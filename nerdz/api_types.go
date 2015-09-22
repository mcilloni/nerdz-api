package nerdz

import (
	"database/sql"
	"time"
)

// Response represent the response format of the API
type Response struct {
	Data         interface{} `json:"data"`
	Message      string      `json:"message"`
	HumanMessage string      `json:"humanMessage"`
	Status       uint        `json:"status"`
	Success      bool        `json:"success"`
}

//Renderable represents a common interface for all the data that
//will be generated by the backend and will be returned by the API
type Renderable interface {
	//Render returns only an identifier of the original data structure
	Render() string
}

// UserPostsNoNotifyTO represents the TO of UserPostsNoNotify
type UserPostsNoNotifyTO struct {
	User    uint64    `json:"user"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (u UserPostsNoNotifyTO) Render() string {
	return "UserPostsNoNotify"
}

// UserPostCommentsNoNotifyTO represents the TO of UserPostCommentsNoNotify
type UserPostCommentsNoNotifyTO struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (u UserPostCommentsNoNotifyTO) Render() string {
	return "UserPostCommentsNoNotify"
}

// UserPostCommentsNotifyTO  represents the TO of UserPostCommentsNotify
type UserPostCommentsNotifyTO struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (u UserPostCommentsNotifyTO) Render() string {
	return "UserPostCommentsNotify"
}

// BanTO represents the TO of Ban
type BanTO struct {
	User       uint64    `json:"user"`
	Motivation string    `json:"motivation"`
	Time       time.Time `json:"time"`
	Counter    uint64    `json:"counter"`
}

func (BanTO) Render() string {
	return "Ban"
}

// BlacklistTO represens the TO of Blacklist
type BlacklistTO struct {
	From       uint64    `json:"from"`
	To         uint64    `json:"to"`
	Motivation string    `json:"motivation"`
	Time       time.Time `json:"time"`
	Counter    uint64    `json:"counter"`
}

func (BlacklistTO) Render() string {
	return "Blacklist"
}

// WhitelistTO represents the TO of Whitelist
type WhitelistTO struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (WhitelistTO) Render() string {
	return "WhiteList"
}

//UserFollowerTO represents the TO of UserFollower
type UserFollowerTO struct {
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `json:"counter"`
}

func (UserFollowerTO) Render() string {
	return "UserFollower"
}

// ProjectNotifyTO represents the TO of ProjectNotify
type ProjectNotifyTO struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Time    time.Time `json:"time"`
	Hpid    uint64    `json:"hpid"`
	Counter uint64    `json:"counter"`
}

func (ProjectNotifyTO) Render() string {
	return "ProjectNotify"
}

// ProjectPostsNoNotifyTo represents the TO of ProjectPostsNoNotify
type ProjectPostsNoNotifyTO struct {
	User    uint64    `json:"user"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostsNoNotifyTO) Render() string {
	return "ProjectPostsNoNotify"
}

// ProjectPostCommentsNoNotifyTO represents the TO of ProjectPostCommentsNoNotify
type ProjectPostCommentsNoNotifyTO struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostCommentsNoNotifyTO) Render() string {
	return "ProjectPostCommentsNoNotify"
}

// ProjectPostCommentsNotifyTO represents the TO of ProjectPostCommentsNotify
type ProjectPostCommentsNotifyTO struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostCommentsNotifyTO) Render() string {
	return "ProjectPostCommentsNotifyTO"
}

// UserTO represents the TO of User
type UserTO struct {
	Counter          uint64    `json:"counter"`
	Last             time.Time `json:"last"`
	NotifyStory      []byte    `json:"notifyStory"`
	Private          bool      `json:"private"`
	Lang             string    `json:"lang"`
	Username         string    `json:"username"`
	Name             string    `json:"name"`
	Surname          string    `json:"surname"`
	Gender           bool      `json:"gender"`
	BirthDate        time.Time `json:"birthDate"`
	BoardLang        string    `json:"boardLang"`
	Timezone         string    `json:"timezone"`
	Viewonline       bool      `json:"viewonline"`
	RegistrationTime time.Time `json:"registrationTime"`
	Profile          Profile
}

func (UserTO) Render() string {
	return "User"
}

// ProfileTO represents the TO of Profile
type ProfileTO struct {
	Counter        uint64    `json:"counter"`
	Website        string    `json:"website"`
	Quotes         string    `json:"quotes"`
	Biography      string    `json:"biography"`
	Interests      string    `json:"interests"`
	Github         string    `json:"github"`
	Skype          string    `json:"skype"`
	Jabber         string    `json:"jabber"`
	Yahoo          string    `json:"yahoo"`
	Userscript     string    `json:"userscript"`     // ?API?
	Template       uint8     `json:"template"`       // ?API?
	MobileTemplate uint8     `json:"mobileTemplate"` // ?API?
	Dateformat     string    `json:"dateformat"`     // ?API?
	Facebook       string    `json:"facebook"`
	Twitter        string    `json:"twitter"`
	Steam          string    `json:"steam"`
	Push           bool      `json:"push"`        // ?API?
	Pushregtime    time.Time `json:"pushregtime"` // ?API?
	Closed         bool      `json:"closed"`
}

func (ProfileTO) Render() string {
	return "Profile"
}

// PostFields represents the common fields presents in a Post
type PostFields struct {
	FromInfo         *Info  `json:"from"`
	ToInfo           *Info  `json:"to"`
	Rate             int    `json:"rate"`
	RevisionsCount   uint8  `json:"revisions"`
	CommentsCount    uint8  `json:"comments"`
	BookmarkersCount uint8  `json:"bookmarkers"`
	LurkersCount     uint8  `json:"lurkers"`
	URL              string `json:"url"`
	Timestamp        int64  `json:"timestamp"`
	CanEdit          bool   `json:"canEdit"`
	CanDelete        bool   `json:"canDelete"`
	CanComment       bool   `json:"canComment"`
	CanBookmark      bool   `json:"canBookmark"`
	CanLurk          bool   `json:"canLurk"`
}

// UserPostTO represents the TO of UserPost
type UserPostTO struct {
	Hpid     uint64    `json:"hpid"`
	Pid      uint64    `json:"pid"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	Lang     string    `json:"lang"`
	News     bool      `json:"news"`
	Closed   bool      `json:"closed"`
	PostInfo PostFields
}

//SetPostFields populate the API fileds of the UserPost
func (post *UserPostTO) SetPostFields(user *User, currPost *UserPost) {
	if from, e := NewUser(currPost.From); e == nil {
		post.PostInfo.FromInfo = from.Info()
	}
	if to, e := NewUser(currPost.To); e == nil {
		post.PostInfo.ToInfo = to.Info()
	}

	post.PostInfo.Rate = currPost.Thumbs()
	post.PostInfo.RevisionsCount = currPost.RevisionsNumber()
	post.PostInfo.CommentsCount = currPost.CommentsNumber()
	post.PostInfo.BookmarkersCount = currPost.BookmarkersNumber()
	post.PostInfo.LurkersCount = currPost.LurkersNumber()
	post.PostInfo.Timestamp = currPost.Time.Unix()
	post.PostInfo.URL = currPost.URL(Configuration.NERDZURL).String()
	post.PostInfo.CanBookmark = user.canBookmark(currPost)
	post.PostInfo.CanComment = user.canComment(currPost)
	post.PostInfo.CanDelete = user.canDelete(currPost)
	post.PostInfo.CanEdit = user.canEdit(currPost)
	post.PostInfo.CanLurk = user.canLurk(currPost)

}

func (UserPostTO) Render() string {
	return "UserPost"
}

// UserPostRevisionTO represents the TO of UserPostRevision
type UserPostRevisionTO struct {
	Hpid    uint64    `json:"hpid"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	RevNo   uint16    `json:"revNo"`
	Counter uint64    `json:"counter"`
}

func (UserPostRevisionTO) Render() string {
	return "UserPostRevision"
}

// UserPostThumbTO represents the TO of UserPostThumb
type UserPostThumbTO struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Vote    int8      `json:"vote"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (UserPostThumbTO) Render() string {
	return "UserPostThumb"
}

// UserPostLurkerTO represents the TO of UserPostLurker
type UserPostLurkerTO struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (UserPostLurkerTO) Render() string {
	return "UserPostLurker"
}

// UserPostCommentTO represents the TO of UserPostComment
type UserPostCommentTO struct {
	Hcid     uint64    `json:"hcid"`
	Hpid     uint64    `json:"hpid"`
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	Editable bool      `json:"editable"`
}

func (UserPostCommentTO) Render() string {
	return "UserPostComment"
}

// UserPostCommentRevisionTO represents the TO of UserPostCommentRevision
type UserPostCommentRevisionTO struct {
	Hcid    uint64    `json:"hcid"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	RevNo   int8      `json:"revNo"`
	Counter uint64    `json:"counter"`
}

func (UserPostCommentRevisionTO) Render() string {
	return "UserPostCommentRevision"
}

// UserPostBookmarkTO represents the TO of UserPostBookmark
type UserPostBookmarkTO struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (UserPostBookmarkTO) Render() string {
	return "UserPostBookmarkTO"
}

// PmTO represents the TO of Pm
type PmTO struct {
	Pmid    uint64    `json:"pmid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Message string    `json:"message"`
	ToRead  bool      `json:"toRead"`
	Time    time.Time `json:"time"`
}

func (PmTO) Render() string {
	return "Pm"
}

// ProjectTO represents the TO of Project
type ProjectTO struct {
	Counter      uint64         `json:"counter"`
	Description  string         `json:"description"`
	Name         string         `json:"name"`
	Private      bool           `json:"private"`
	Photo        sql.NullString `json:"photo"`
	Website      sql.NullString `json:"website"`
	Goal         string         `json:"goal"`
	Visible      bool           `json:"visible"`
	Open         bool           `json:"open"`
	CreationTime time.Time      `json:"creationTime"`
}

func (ProjectTO) Render() string {
	return "Project"
}

// ProjectMemberTO represents the TO of ProjectMember
type ProjectMemberTO struct {
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `json:"counter"`
}

func (ProjectMemberTO) Render() string {
	return "ProjectMember"
}

// ProjectOwnerTO represents the TO of ProjectOwner
type ProjectOwnerTO struct {
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `json:"counter"`
}

func (ProjectOwnerTO) Render() string {
	return "ProjectOwner"
}

// ProjectPostTO represents the TO of ProjectPost
type ProjectPostTO struct {
	Hpid     uint64    `json:"hpid"`
	Pid      uint64    `json:"pid"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	News     bool      `json:"news"`
	Lang     string    `json:"lang"`
	Closed   bool      `json:"closed"`
	PostInfo PostFields
}

//SetPostFields populate the API fileds of the UserPost
func (post *ProjectPostTO) SetPostFields(user *User, currPost *ProjectPost) {
	if from, e := NewUser(currPost.From); e == nil {
		post.PostInfo.FromInfo = from.Info()
	}
	if to, e := NewProject(currPost.To); e == nil {
		post.PostInfo.ToInfo = to.Info()
	}

	post.PostInfo.Rate = currPost.Thumbs()
	post.PostInfo.RevisionsCount = currPost.RevisionsNumber()
	post.PostInfo.CommentsCount = currPost.CommentsNumber()
	post.PostInfo.BookmarkersCount = currPost.BookmarkersNumber()
	post.PostInfo.LurkersCount = currPost.LurkersNumber()
	post.PostInfo.Timestamp = currPost.Time.Unix()
	post.PostInfo.URL = currPost.URL(Configuration.NERDZURL).String()
	post.PostInfo.CanBookmark = user.canBookmark(currPost)
	post.PostInfo.CanComment = user.canComment(currPost)
	post.PostInfo.CanDelete = user.canDelete(currPost)
	post.PostInfo.CanEdit = user.canEdit(currPost)
	post.PostInfo.CanLurk = user.canLurk(currPost)
}

func (ProjectPostTO) Render() string {
	return "ProjectPost"
}

// ProjectPostRevisionTO represents the TO of ProjectPostRevision
type ProjectPostRevisionTO struct {
	Hpid    uint64    `json:"hpid"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	RevNo   uint16    `json:"revNo"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostRevisionTO) Render() string {
	return "ProjectPost"
}

// ProjectPostThumbTO represents the TO of ProjectPostThumb
type ProjectPostThumbTO struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Time    time.Time `json:"time"`
	Vote    int8      `json:"vote"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostThumbTO) Render() string {
	return "ProjectPostThumb"
}

// ProjectPostLurkerTO represents the TO of ProjectPostLurker
type ProjectPostLurkerTO struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostLurkerTO) Render() string {
	return "ProjectPostLurker"
}

// ProjectPostCommentTO represents the TO of ProjectPostComment
type ProjectPostCommentTO struct {
	Hcid     uint64    `json:"hcid"`
	Hpid     uint64    `json:"hpid"`
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	Editable bool      `json:"editable"`
}

func (ProjectPostCommentTO) Render() string {
	return "ProjectPostComment"
}

// ProjectPostCommentRevisionTO represents the TO of ProjectPostCommentRevision
type ProjectPostCommentRevisionTO struct {
	Hcid    uint64    `json:"hcid"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	RevNo   uint16    `json:"revNo"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostCommentRevisionTO) Render() string {
	return "ProjectPostCommentRevision"
}

// ProjectPostBookmarkTO represents the TO of ProjectPostBookmark
type ProjectPostBookmarkTO struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostBookmarkTO) Render() string {
	return "ProjectPostBookmark"
}

// ProjectFollowerTO represents the TO of ProjectFollower
type ProjectFollowerTO struct {
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `json:"counter"`
}

func (ProjectFollowerTO) Render() string {
	return "ProjectFollower"
}

// UserPostCommentThumbTO represents the TO of UserPostCommentThumb
type UserPostCommentThumbTO struct {
	Hcid    uint64 `json:"hcid"`
	User    uint64 `json:"user"`
	Vote    int8   `json:"vote"`
	Counter uint64 `json:"counter"`
}

func (UserPostCommentThumbTO) Render() string {
	return "UserPostCommentThumb"
}

// ProjectPostCommentThumbTO represents the TO of ProjectPostCommentThumb
type ProjectPostCommentThumbTO struct {
	Hcid    uint64    `json:"hcid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Vote    int8      `json:"vote"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostCommentThumbTO) Render() string {
	return "ProjectPostCommentThumb"
}

// DeletedUserTO represents the TO of DeletedUserTO
type DeletedUserTO struct {
	Counter    uint64    `gorm:"primary_key:yes" json:"counter"`
	Username   string    `json:"username"`
	Time       time.Time `sql:"default:NOW()" json:"time"`
	Motivation string    `json:"motivation"`
}

func (DeletedUserTO) Render() string {
	return "DeletedUser"
}

// SpecialUserTO represents the TO of SpecialUser
type SpecialUserTO struct {
	Role    string `json:"role"`
	Counter uint64 `json:"counter"`
}

func (SpecialUserTO) Render() string {
	return "SpecialUser"
}

// SpecialProjectTO represents the TO of SpecialProject
type SpecialProjectTO struct {
	Role    string `json:"role"`
	Counter uint64 `json:"counter"`
}

func (SpecialProjectTO) Render() string {
	return "SpecialProject"
}

// PostClassificationTO represents the TO of PostClassification
type PostClassificationTO struct {
	ID    uint64 `json:"id"`
	UHpid uint64 `json:"uHpid"`
	GHpid uint64 `json:"gHpid"`
	Tag   string `json:"tag"`
}

func (PostClassificationTO) Render() string {
	return "PostClassification"
}

// MentionTO represents the TO of Mention
type MentionTO struct {
	ID       uint64    `json:"id"`
	UHpid    uint64    `json:"uHpid"`
	GHpid    uint64    `json:"gHpid"`
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
}

func (MentionTO) Render() string {
	return "Mention"
}

type PersonalInfoTO struct {
	ID        uint64    `json:"id"`
	IsOnline  bool      `json:"online"`
	Nation    string    `json:"nation"`
	Timezone  string    `json:"timezone"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Gender    bool      `json:"gender"`
	Birthday  time.Time `json:"birthday"`
	Gravatar  string    `json:"gravatar"`
	Interests []string  `json:"interests"`
	Quotes    []string  `json:"quotes"`
	Biography string    `json:"biography"`
}

func (PersonalInfoTO) Render() string {
	return "PersonalInfo"
}

// ContactInfoTO represents the TO of ContactInfo
type ContactInfoTO struct {
	Website  string `json:"website"`
	GitHub   string `json:"github"`
	Skype    string `json:"skype"`
	Jabber   string `json:"jabber"`
	Yahoo    string `json:"yahoo"`
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
	Steam    string `json:"steam"`
}

func (ContactInfoTO) Render() string {
	return "ContactInfo"
}

// InfoTO represents the TO of Info
type InfoTO struct {
	ID          uint64     `json:"id"`
	Owner       Renderable `json:"owner"`
	Name        string     `json:"name"`
	Username    string     `json:"username"`
	Website     string     `json:"website"`
	Image       string     `json:"image"`
	Closed      bool       `json:"closed"`
	Type        boardType  `json:"type"`
	BoardString string     `json:"board"`
}

func (InfoTO) Render() string {
	return "Info"
}
