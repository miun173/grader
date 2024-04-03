// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: course_assignments.reader.sql

package xsqlc

import (
	"context"
)

const findAllAssignmentsByCourseID = `-- name: FindAllAssignmentsByCourseID :many
SELECT id, assigned_by, name, description, case_output_file_id, case_input_file_id, created_at, deadline_at, updated_at, deleted_at, template FROM assignments WHERE id IN (
    SELECT assignment_id FROM rel_assignment_to_course WHERE course_id = $1
)
ORDER BY updated_at DESC
LIMIT $3
OFFSET $2
`

type FindAllAssignmentsByCourseIDParams struct {
	CourseID   string
	PageOffset int32
	PageLimit  int32
}

func (q *Queries) FindAllAssignmentsByCourseID(ctx context.Context, arg FindAllAssignmentsByCourseIDParams) ([]Assignment, error) {
	rows, err := q.db.QueryContext(ctx, findAllAssignmentsByCourseID, arg.CourseID, arg.PageOffset, arg.PageLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Assignment
	for rows.Next() {
		var i Assignment
		if err := rows.Scan(
			&i.ID,
			&i.AssignedBy,
			&i.Name,
			&i.Description,
			&i.CaseOutputFileID,
			&i.CaseInputFileID,
			&i.CreatedAt,
			&i.DeadlineAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Template,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
