// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Node interface {
	IsNode()
}

type EditNoteInput struct {
	Title *string     `json:"title"`
	Text  *string     `json:"text"`
	Tags  []*TagInput `json:"tags"`
}

type EditRenderingInput struct {
	Name  *string      `json:"name"`
	Lanes []*LaneInput `json:"lanes"`
}

type Lane struct {
	ID     string `json:"id"`
	Filter string `json:"filter"`
}

func (Lane) IsNode() {}

type LaneInput struct {
	ID     string `json:"id"`
	Filter string `json:"filter"`
}

type Note struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Tags  []*Tag `json:"tags"`
}

func (Note) IsNode() {}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

type Rendering struct {
	ID    string  `json:"id"`
	Name  *string `json:"name"`
	Lanes []*Lane `json:"lanes"`
}

func (Rendering) IsNode() {}

type Tag struct {
	ID string `json:"id"`
}

func (Tag) IsNode() {}

type TagInput struct {
	ID string `json:"id"`
}

type User struct {
	ID                   string                    `json:"id"`
	Name                 string                    `json:"name"`
	Email                string                    `json:"email"`
	Node                 Node                      `json:"node"`
	NotesConnection      *UserNotesConnection      `json:"notesConnection"`
	RenderingsConnection *UserRenderingsConnection `json:"renderingsConnection"`
}

func (User) IsNode() {}

type UserNoteEdge struct {
	Cursor string `json:"cursor"`
	Node   *Note  `json:"node"`
}

type UserNotesConnection struct {
	Edges      []*UserNoteEdge `json:"edges"`
	PageInfo   *PageInfo       `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}

type UserRenderingEdge struct {
	Cursor string     `json:"cursor"`
	Node   *Rendering `json:"node"`
}

type UserRenderingsConnection struct {
	Edges      []*UserRenderingEdge `json:"edges"`
	PageInfo   *PageInfo            `json:"pageInfo"`
	TotalCount int                  `json:"totalCount"`
}
