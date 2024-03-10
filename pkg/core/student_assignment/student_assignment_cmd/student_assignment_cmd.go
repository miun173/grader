package student_assignment_cmd

import (
	"context"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/fahmifan/autograd/pkg/core"
	"github.com/fahmifan/autograd/pkg/core/auth"
	"github.com/fahmifan/autograd/pkg/core/mediastore"
	"github.com/fahmifan/autograd/pkg/core/mediastore/mediastore_cmd"
	"github.com/fahmifan/autograd/pkg/core/student_assignment"
	"github.com/fahmifan/autograd/pkg/dbmodel"
	"github.com/fahmifan/autograd/pkg/logs"
	autogradv1 "github.com/fahmifan/autograd/pkg/pb/autograd/v1"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentAssignmentCmd struct {
	*core.Ctx
}

func (cmd *StudentAssignmentCmd) SubmitStudentSubmission(ctx context.Context, req *connect.Request[autogradv1.SubmitStudentSubmissionRequest]) (
	*connect.Response[autogradv1.CreatedResponse], error,
) {
	authUser, ok := auth.GetUserFromCtx(ctx)
	if !ok {
		return nil, core.ErrUnauthenticated
	}

	if !authUser.Role.Can(auth.CreateSubmission) {
		return nil, core.ErrPermissionDenied
	}

	assignmentID, err := uuid.Parse(req.Msg.GetAssignmentId())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	submissionFileReader := student_assignment.SubmissionFileReader{}
	assignmentReader := student_assignment.StudentAssignmentReader{}
	submissionWriter := student_assignment.StudentSubmissionWriter{}
	mediastoreCmd := &mediastore_cmd.MediaStoreCmd{Ctx: cmd.Ctx}

	studentID := authUser.UserID
	now := time.Now()
	newID := uuid.New()

	err = core.Transaction(ctx, cmd.Ctx, func(tx *gorm.DB) error {
		assignmet, err := assignmentReader.FindByID(ctx, tx, student_assignment.FindStudentAssignmentByIDRequest{
			ID:            assignmentID,
			StudentID:     studentID,
			WithStudentID: true,
		})
		if err != nil {
			logs.ErrCtx(ctx, err, "StudentAssignmentCmd: CreateStudentSubmission: find assignment by id")
			return core.ErrInternalServer
		}

		mediaRes, err := mediastoreCmd.InternalSave(ctx, tx, mediastore_cmd.InternalSaveRequest{
			Ext:       ".cpp",
			Body:      strings.NewReader(req.Msg.GetSubmissionCode()),
			MediaType: mediastore.MediaFileType(dbmodel.FileTypeSubmission),
		})
		if err != nil {
			logs.ErrCtx(ctx, err, "StudentAssignmentCmd: CreateStudentSubmission: save submission code")
			return core.ErrInternalServer
		}

		submissionFile, err := submissionFileReader.FindByID(ctx, tx, mediaRes.ID)
		if err != nil {
			logs.ErrCtx(ctx, err, "StudentAssignmentCmd: CreateStudentSubmission: find submission file by id")
			return core.ErrInternalServer
		}

		submission, err := student_assignment.CreateStudentSubmission(student_assignment.CreateStudentSubmissionRequest{
			NewID: newID,
			Now:   now,
			Student: student_assignment.Student{
				ID:     studentID,
				Name:   authUser.Name,
				Active: true,
			},
			Assignment: student_assignment.Assignment{
				ID:            assignmet.ID,
				DeadlineAt:    assignmet.DeadlineAt,
				HasAssignment: assignmet.HasSubmission,
			},
			SubmissionFile: submissionFile,
		})
		if err != nil {
			return connect.NewError(connect.CodeInvalidArgument, err)
		}

		err = submissionWriter.CreateSubmission(ctx, tx, &submission)
		if err != nil {
			logs.ErrCtx(ctx, err, "StudentAssignmentCmd: CreateStudentSubmission: create student submission")
			return core.ErrInternalServer
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &connect.Response[autogradv1.CreatedResponse]{
		Msg: &autogradv1.CreatedResponse{
			Id: newID.String(),
		},
	}, nil
}

func (cmd *StudentAssignmentCmd) UpdateStudentSubmission(ctx context.Context, req *connect.Request[autogradv1.UpdateStudentSubmissionRequest]) (
	*connect.Response[autogradv1.Empty], error,
) {
	authUser, ok := auth.GetUserFromCtx(ctx)
	if !ok {
		return nil, core.ErrUnauthenticated
	}

	if !authUser.Role.Can(auth.CreateSubmission) {
		return nil, core.ErrPermissionDenied
	}

	submissionID, err := uuid.Parse(req.Msg.GetSubmissionId())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	submissionFileReader := student_assignment.SubmissionFileReader{}
	submissionReader := student_assignment.StudentSubmissionReader{}
	submissionWriter := student_assignment.StudentSubmissionWriter{}
	mediastoreCmd := &mediastore_cmd.MediaStoreCmd{Ctx: cmd.Ctx}

	now := time.Now()

	err = core.Transaction(ctx, cmd.Ctx, func(tx *gorm.DB) error {
		submission, err := submissionReader.FindByID(ctx, tx, submissionID)
		if err != nil {
			logs.ErrCtx(ctx, err, "StudentAssignmentCmd: CreateStudentSubmission: find submission by id")
			return core.ErrInternalServer
		}

		mediaRes, err := mediastoreCmd.InternalSave(ctx, tx, mediastore_cmd.InternalSaveRequest{
			Ext:       ".cpp",
			Body:      strings.NewReader(req.Msg.GetSubmissionCode()),
			MediaType: mediastore.MediaFileType(dbmodel.FileTypeSubmission),
		})
		if err != nil {
			logs.ErrCtx(ctx, err, "StudentAssignmentCmd: CreateStudentSubmission: save submission code")
			return core.ErrInternalServer
		}

		submissionFile, err := submissionFileReader.FindByID(ctx, tx, mediaRes.ID)
		if err != nil {
			logs.ErrCtx(ctx, err, "StudentAssignmentCmd: CreateStudentSubmission: find submission file by id")
			return core.ErrInternalServer
		}

		submission, err = submission.Update(student_assignment.UpdateStudentSubmissionRequest{
			Now:               now,
			NewSubmissionFile: submissionFile,
		})
		if err != nil {
			return connect.NewError(connect.CodeInvalidArgument, err)
		}

		err = submissionWriter.UpdateSubmission(ctx, tx, &submission)
		if err != nil {
			logs.ErrCtx(ctx, err, "StudentAssignmentCmd: CreateStudentSubmission: create student submission")
			return core.ErrInternalServer
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return core.ProtoEmptyResponse, nil
}
