package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.53

import (
	"context"
	"fmt"
	"log"

	"github.com/99designs/gqlgen/graphql"
	gqlgen "github.com/BRAVO68WEB/collaborate-with-me/packages/api/graphql"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/graphql/model"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/repository"
	appContext "github.com/BRAVO68WEB/collaborate-with-me/packages/api/utils"
)

// CreateWorkspace is the resolver for the createWorkspace field.
func (r *mutationResolver) CreateWorkspace(ctx context.Context, input model.NewWorkspace) (*model.Workspace, error) {
	panic(fmt.Errorf("not implemented: CreateWorkspace - createWorkspace"))
}

// UpdateWorkspace is the resolver for the updateWorkspace field.
func (r *mutationResolver) UpdateWorkspace(ctx context.Context, id string, input model.NewWorkspace) (*model.Workspace, error) {
	panic(fmt.Errorf("not implemented: UpdateWorkspace - updateWorkspace"))
}

// DeleteWorkspace is the resolver for the deleteWorkspace field.
func (r *mutationResolver) DeleteWorkspace(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteWorkspace - deleteWorkspace"))
}

// AddUserToWorkspace is the resolver for the addUserToWorkspace field.
func (r *mutationResolver) AddUserToWorkspace(ctx context.Context, workspaceID string, userID string) (*model.Workspace, error) {
	panic(fmt.Errorf("not implemented: AddUserToWorkspace - addUserToWorkspace"))
}

// RemoveUserFromWorkspace is the resolver for the removeUserFromWorkspace field.
func (r *mutationResolver) RemoveUserFromWorkspace(ctx context.Context, workspaceID string, userID string) (*model.Workspace, error) {
	panic(fmt.Errorf("not implemented: RemoveUserFromWorkspace - removeUserFromWorkspace"))
}

// AddExcalidrawObject is the resolver for the addExcalidrawObject field.
func (r *mutationResolver) AddExcalidrawObject(ctx context.Context, workspaceID string, object any) (*model.Workspace, error) {
	panic(fmt.Errorf("not implemented: AddExcalidrawObject - addExcalidrawObject"))
}

// RemoveExcalidrawObject is the resolver for the removeExcalidrawObject field.
func (r *mutationResolver) RemoveExcalidrawObject(ctx context.Context, workspaceID string, objectID string) (*model.Workspace, error) {
	panic(fmt.Errorf("not implemented: RemoveExcalidrawObject - removeExcalidrawObject"))
}

// SingleUpload is the resolver for the singleUpload field.
func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (*model.UploadResponse, error) {
	panic(fmt.Errorf("not implemented: SingleUpload - singleUpload"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	result, err := r.Domain.User.CreateUser(input.Email, input.Password, input.Username)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &model.User{
		ID:        result.ID.Hex(),
		Email:     result.Email,
		Username:  result.Username,
		IsActive:  result.IsActive,
		Role:      result.Role,
		CreatedAt: result.CreatedAt.Time().GoString(),
		UpdatedAt: result.UpdatedAt.Time().GoString(),
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdateUser) (*model.User, error) {
	userID, err := appContext.UserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	is_admin := r.Domain.User.CheckIfUserIsAdmin(userID)

	if userID != id && !is_admin {
		return nil, fmt.Errorf("access denied")
	}

	updateUser := repository.UpdateUser{}

	if input.Username != nil && *input.Username != "" {
		updateUser.Username = *input.Username
	}

	if input.Password != nil && *input.Password != "" {
		updateUser.Password = *input.Password
	}

	if input.Role != nil && *input.Role != "" {
		if !is_admin {
			return nil, fmt.Errorf("access denied")
		}

		updateUser.Role = *input.Role
	}

	if input.Email != nil && *input.Email != "" {
		if !is_admin {
			return nil, fmt.Errorf("access denied")
		}

		updateUser.Email = *input.Email
	}

	result, err := r.Domain.User.UpdateUserByID(id, updateUser)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &model.User{
		ID:        result.ID.Hex(),
		Email:     result.Email,
		Username:  result.Username,
		IsActive:  result.IsActive,
		Role:      result.Role,
		CreatedAt: result.CreatedAt.Time().GoString(),
		UpdatedAt: result.UpdatedAt.Time().GoString(),
	}, nil
}

// DisableUser is the resolver for the disableUser field.
func (r *mutationResolver) DisableUser(ctx context.Context, id string) (bool, error) {
	userID, err := appContext.UserIDFromContext(ctx)

	if err != nil {
		return false, err
	}

	is_admin := r.Domain.User.CheckIfUserIsAdmin(userID)

	if !is_admin {
		return false, fmt.Errorf("access denied")
	}

	_, err = r.Domain.User.DisableUserByID(id)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.LoginResponse, error) {
	token, err := r.Domain.User.Login(email, password)

	if err != nil {
		return &model.LoginResponse{
			IsSuccess:   false,
			AccessToken: "",
		}, err
	}

	return &model.LoginResponse{
		IsSuccess:   true,
		AccessToken: token,
	}, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	_, err := appContext.UserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	users, err := r.Domain.User.GetUsers(
		1,
	)

	if err != nil {
		return nil, err
	}

	var result []*model.User

	for _, user := range users {
		result = append(result, &model.User{
			ID:        user.ID.Hex(),
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Role,
			IsActive:  user.IsActive,
			CreatedAt: user.CreatedAt.Time().GoString(),
			UpdatedAt: user.UpdatedAt.Time().GoString(),
		})
	}

	return result, nil
}

// Workspaces is the resolver for the workspaces field.
func (r *queryResolver) Workspaces(ctx context.Context, userID *string) ([]*model.Workspace, error) {
	panic(fmt.Errorf("not implemented: Workspaces - workspaces"))
}

// Workspace is the resolver for the workspace field.
func (r *queryResolver) Workspace(ctx context.Context, id string, userID string) (*model.Workspace, error) {
	panic(fmt.Errorf("not implemented: Workspace - workspace"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	_, err := appContext.UserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	user, err := r.Domain.User.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        user.ID.Hex(),
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt.Time().GoString(),
		UpdatedAt: user.UpdatedAt.Time().GoString(),
	}, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	userID, err := appContext.UserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	user, err := r.Domain.User.GetUserByID(userID)

	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        user.ID.Hex(),
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt.Time().GoString(),
		UpdatedAt: user.UpdatedAt.Time().GoString(),
	}, nil
}

// LiveWorkspaceUpdates is the resolver for the liveWorkspaceUpdates field.
func (r *subscriptionResolver) LiveWorkspaceUpdates(ctx context.Context, workspaceID string) (<-chan any, error) {
	panic(fmt.Errorf("not implemented: LiveWorkspaceUpdates - liveWorkspaceUpdates"))
}

// LiveUserUpdates is the resolver for the liveUserUpdates field.
func (r *subscriptionResolver) LiveUserUpdates(ctx context.Context, userID string) (<-chan *model.User, error) {
	panic(fmt.Errorf("not implemented: LiveUserUpdates - liveUserUpdates"))
}

// LiveWorkspaceCollaborators is the resolver for the liveWorkspaceCollaborators field.
func (r *subscriptionResolver) LiveWorkspaceCollaborators(ctx context.Context, workspaceID string) (<-chan []*model.User, error) {
	panic(fmt.Errorf("not implemented: LiveWorkspaceCollaborators - liveWorkspaceCollaborators"))
}

// Mutation returns gqlgen.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgen.MutationResolver { return &mutationResolver{r} }

// Query returns gqlgen.QueryResolver implementation.
func (r *Resolver) Query() gqlgen.QueryResolver { return &queryResolver{r} }

// Subscription returns gqlgen.SubscriptionResolver implementation.
func (r *Resolver) Subscription() gqlgen.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
