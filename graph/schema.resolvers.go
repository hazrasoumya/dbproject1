package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sayani/dbproject1/graph/generated"
	"github.com/sayani/dbproject1/graph/model"
	"github.com/sayani/dbproject1/graph/postgres"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVal(ctx context.Context, input model.Inputstr2) (*model.Outputstr2, error) {
	//panic(fmt.Errorf("not implemented"))
	postgres.Check2(input.Roll, input.Name)
	s := model.Outputstr2{"abcd"}
	return &s, nil
}

func (r *mutationResolver) UpdateMarks(ctx context.Context, input model.Marks) (*model.Total, error) {
	//query := `update my_courses set sem1 =$1, sem2 = $2, sem3 =$3 where roll = $4 and course = $5`
	//panic(fmt.Errorf("not implemented"))
	//fmt.Println(input.X, input.Y, input.Z, input.Ro, input.C)
	postgres.Calculate(input.X, input.Y, input.Z, input.Ro, input.C)
	// var i int
	// postgres.Test()
	s := model.Total{1}
	return &s, nil
}

func (r *mutationResolver) PutCourse(ctx context.Context, input model.Inputcourses) (*model.Outpurcourses, error) {
	//panic(fmt.Errorf("not implemented"))
	coursestr := postgres.Check01(input.X, input.Y)
	var str []*string
	for i := range coursestr {
		str = append(str, &coursestr[i])
	}
	s := model.Outpurcourses{str}
	return &s, nil
}

func (r *mutationResolver) UpdateCourse(ctx context.Context, input model.Inputcourses) (*model.Outpurcourses, error) {
	//panic(fmt.Errorf("not implemented"))
	coursestr := postgres.Check11(input.X, input.Y)
	// strr := "sayani"
	// strr1 := "das"
	// var str1 []*string
	// str1 = append(str1, &strr)
	// str1 = append(str1, &strr1)
	// fmt.Println("my", str1)
	var str []*string
	for i, ele := range coursestr {
		fmt.Println(i, ele)
		str = append(str, &ele)
		//	fmt.Println("mine", str)
	}
	s := model.Outpurcourses{str}
	return &s, nil
}

func (r *mutationResolver) DeleteCourse(ctx context.Context, input model.Inputcourses) (*model.Outpurcourses, error) {
	//panic(fmt.Errorf("not implemented"))
	postgres.Check21(input.X, input.Y)
	var str []*string
	/*for i := range coursestr {
		str = append(str, &coursestr[i])
	}*/
	s := model.Outpurcourses{str}
	return &s, nil
}

func (r *mutationResolver) Testcase(ctx context.Context, input *model.Testmarks) (*model.Testtotal, error) {
	//panic(fmt.Errorf("not implemented"))
	postgres.Testfunc(input.X)
	i := 2
	s := model.Testtotal{i}
	return &s, nil
}

func (r *mutationResolver) InsertVal(ctx context.Context, input model.Inputstr) (*model.Outputstr, error) {
	//	panic(fmt.Errorf("not implemented"))
	p := input.X
	q := input.Y
	st := input.P
	postgres.Check(p, q, st)
	s := model.Outputstr{"abcd"}
	return &s, nil
}

func (r *mutationResolver) UpdateVal(ctx context.Context, input model.Inputstr1) (*model.Outputstr1, error) {
	//panic(fmt.Errorf("not implemented"))
	postgres.Check1(input.Roll, input.Course, input.Sem1, input.Sem2, input.Sem3)
	s := model.Outputstr1{11}
	return &s, nil
}

func (r *mutationResolver) UpdateNameroll(ctx context.Context, input model.Inputstr2) (*model.Outputstr2, error) {
	//	panic(fmt.Errorf("not implemented"))
	str := postgres.Updatename(input.Roll, input.Name)
	s := model.Outputstr2{str}
	return &s, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DisplayVal(ctx context.Context, input model.Inputvalues) (*model.Outputvalues, error) {
	//panic(fmt.Errorf("not implemented"))
	postgres.Check3(input.X, input.Y)
	s := model.Outputvalues{"displayed"}
	return &s, nil
}

func (r *queryResolver) DisplayVal1(ctx context.Context, input model.Inputvalues1) (*model.Outputvalues1, error) {
	//	panic(fmt.Errorf("not implemented"))
	arr := postgres.Check4(input.X, input.Y)
	fmt.Println(arr)
	var arr1 []*string
	//	var arr2 []*string
	//var arr2 []string
	//	arr1 = &arr
	//	arr1=&arr2
	//	fmt.Println(arr1)
	//	s := model.Outputvalues1{arr2}
	//	arr1 = arr
	// i := 0
	for i := range arr {
		/*if s == " " {
			arr1 = append(arr1, &s)
			continue
		}*/
		arr1 = append(arr1, &arr[i])
	}
	s := model.Outputvalues1{arr1}
	return &s, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
