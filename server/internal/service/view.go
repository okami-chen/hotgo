// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/model"
)

type (
	IView interface {
		GetBreadCrumb(ctx context.Context, in *model.ViewGetBreadCrumbInput) []model.ViewBreadCrumb
		GetTitle(ctx context.Context, in *model.ViewGetTitleInput) string
		RenderTpl(ctx context.Context, tpl string, data ...model.View)
		Render(ctx context.Context, data ...model.View)
		Error(ctx context.Context, err error)
	}
)

var (
	localView IView
)

func View() IView {
	if localView == nil {
		panic("implement not found for interface IView, forgot register?")
	}
	return localView
}

func RegisterView(i IView) {
	localView = i
}
