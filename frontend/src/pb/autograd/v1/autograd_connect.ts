// @generated by protoc-gen-connect-es v0.12.0 with parameter "target=ts"
// @generated from file autograd/v1/autograd.proto (package autograd.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import {
	Assignment,
	CreateAssignmentRequest,
	CreatedResponse,
	CreateManagedUserRequest,
	CreateSubmissionRequest,
	DeleteByIDRequest,
	Empty,
	FindAllAssignmentsRequest,
	FindAllAssignmentsResponse,
	FindAllManagedUsersRequest,
	FindAllManagedUsersResponse,
	FindByIDRequest,
	LoginRequest,
	LoginResponse,
	PingResponse,
	Submission,
	UpdateAssignmentRequest,
	UpdateSubmissionRequest,
} from "./autograd_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service autograd.v1.AutogradService
 */
export const AutogradService = {
	typeName: "autograd.v1.AutogradService",
	methods: {
		/**
		 * @generated from rpc autograd.v1.AutogradService.Ping
		 */
		ping: {
			name: "Ping",
			I: Empty,
			O: PingResponse,
			kind: MethodKind.Unary,
		},
		/**
		 * User Management
		 *
		 * @generated from rpc autograd.v1.AutogradService.CreateManagedUser
		 */
		createManagedUser: {
			name: "CreateManagedUser",
			I: CreateManagedUserRequest,
			O: CreatedResponse,
			kind: MethodKind.Unary,
		},
		/**
		 * @generated from rpc autograd.v1.AutogradService.FindAllManagedUsers
		 */
		findAllManagedUsers: {
			name: "FindAllManagedUsers",
			I: FindAllManagedUsersRequest,
			O: FindAllManagedUsersResponse,
			kind: MethodKind.Unary,
		},
		/**
		 * Assignment Submission
		 * Assignment Queries
		 *
		 * @generated from rpc autograd.v1.AutogradService.FindAssignment
		 */
		findAssignment: {
			name: "FindAssignment",
			I: FindByIDRequest,
			O: Assignment,
			kind: MethodKind.Unary,
		},
		/**
		 * @generated from rpc autograd.v1.AutogradService.FindAllAssignments
		 */
		findAllAssignments: {
			name: "FindAllAssignments",
			I: FindAllAssignmentsRequest,
			O: FindAllAssignmentsResponse,
			kind: MethodKind.Unary,
		},
		/**
		 * @generated from rpc autograd.v1.AutogradService.FindSubmission
		 */
		findSubmission: {
			name: "FindSubmission",
			I: FindByIDRequest,
			O: Submission,
			kind: MethodKind.Unary,
		},
		/**
		 * Assignment Mutations
		 *
		 * @generated from rpc autograd.v1.AutogradService.CreateAssignment
		 */
		createAssignment: {
			name: "CreateAssignment",
			I: CreateAssignmentRequest,
			O: CreatedResponse,
			kind: MethodKind.Unary,
		},
		/**
		 * @generated from rpc autograd.v1.AutogradService.UpdateAssignment
		 */
		updateAssignment: {
			name: "UpdateAssignment",
			I: UpdateAssignmentRequest,
			O: Empty,
			kind: MethodKind.Unary,
		},
		/**
		 * @generated from rpc autograd.v1.AutogradService.DeleteAssignment
		 */
		deleteAssignment: {
			name: "DeleteAssignment",
			I: DeleteByIDRequest,
			O: Empty,
			kind: MethodKind.Unary,
		},
		/**
		 * @generated from rpc autograd.v1.AutogradService.CreateSubmission
		 */
		createSubmission: {
			name: "CreateSubmission",
			I: CreateSubmissionRequest,
			O: CreatedResponse,
			kind: MethodKind.Unary,
		},
		/**
		 * @generated from rpc autograd.v1.AutogradService.UpdateSubmission
		 */
		updateSubmission: {
			name: "UpdateSubmission",
			I: UpdateSubmissionRequest,
			O: Empty,
			kind: MethodKind.Unary,
		},
		/**
		 * @generated from rpc autograd.v1.AutogradService.DeleteSubmission
		 */
		deleteSubmission: {
			name: "DeleteSubmission",
			I: DeleteByIDRequest,
			O: Empty,
			kind: MethodKind.Unary,
		},
		/**
		 * Auth
		 * Auth Queries
		 *
		 * @generated from rpc autograd.v1.AutogradService.Login
		 */
		login: {
			name: "Login",
			I: LoginRequest,
			O: LoginResponse,
			kind: MethodKind.Unary,
		},
	},
} as const;
