package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"backend/graph/generated"
	"backend/graph/model"
	"backend/models"
	"context"
	"errors"
	"strings"

	"github.com/go-pg/pg/v10/orm"
)

func (r *categoryResolver) Videos(ctx context.Context, obj *models.Category) ([]*models.Video, error) {
	var cgv []*models.Video
	var videos []*models.Video
	err := r.DB.Model(&videos).Select()
	if err != nil {
		return nil, errors.New("failed to select video :(")
	}
	for _, video := range videos {
		if video.CategoryID == obj.ID {
			cgv = append(cgv, video)
		}
	}
	return cgv, nil
}

func (r *commentResolver) Replies(ctx context.Context, obj *models.Comment) ([]*models.Reply, error) {
	var cgv []*models.Reply
	var videos []*models.Reply
	err := r.DB.Model(&videos).Select()
	if err != nil {
		return nil, errors.New("failed to select video :(")
	}
	for _, video := range videos {
		if video.CommentID == obj.ID {
			cgv = append(cgv, video)
		}
	}
	return cgv, nil
}

func (r *commentResolver) Creator(ctx context.Context, obj *models.Comment) (*models.Creator, error) {
	var creators []*models.Creator
	var creator *models.Creator
	err := r.DB.Model(&creators).Select()
	if err != nil {
		return nil, errors.New("failed to select creator :(")
	}
	for _, c := range creators {
		if c.ID == obj.CreatorID {
			creator = c
			break
		}
	}
	return creator, nil
}

func (r *communityResolver) Creator(ctx context.Context, obj *models.Community) (*models.Creator, error) {
	var creators []*models.Creator
	var creator *models.Creator
	err := r.DB.Model(&creators).Select()
	if err != nil {
		return nil, errors.New("failed to select creator :(")
	}
	for _, c := range creators {
		if c.ID == obj.CreatorID {
			creator = c
			break
		}
	}
	return creator, nil
}

func (r *creatorResolver) Playlists(ctx context.Context, obj *models.Creator) ([]*models.Playlist, error) {
	var cgv []*models.Playlist
	var videos []*models.Playlist
	err := r.DB.Model(&videos).Select()
	if err != nil {
		return nil, errors.New("failed to select video :(")
	}
	for _, video := range videos {
		if video.CreatorID == obj.ID {
			cgv = append(cgv, video)
		}
	}
	return cgv, nil
}

func (r *creatorResolver) Videos(ctx context.Context, obj *models.Creator) ([]*models.Video, error) {
	var crv []*models.Video
	var videos []*models.Video
	err := r.DB.Model(&videos).Select()
	if err != nil {
		return nil, errors.New("failed to select video :(")
	}
	for _, video := range videos {
		if video.CreatorID == obj.ID {
			crv = append(crv, video)
		}
	}
	return crv, nil
}

func (r *creatorResolver) Subscribers(ctx context.Context, obj *models.Creator) ([]*models.Subscriber, error) {
	var crs []*models.Subscriber
	var subscribers []*models.Subscriber
	err := r.DB.Model(&subscribers).Select()
	if err != nil {
		return nil, errors.New("failed to select subscriber :(")
	}
	for _, subscriber := range subscribers {
		if subscriber.CreatorID == obj.ID {
			crs = append(crs, subscriber)
		}
	}
	return crs, nil
}

func (r *membershipResolver) Creator(ctx context.Context, obj *models.Membership) (*models.Creator, error) {
	var creators []*models.Creator
	var creator *models.Creator
	err := r.DB.Model(&creators).Select()
	if err != nil {
		return nil, errors.New("failed to select creator :(")
	}
	for _, c := range creators {
		if c.ID == obj.CreatorID {
			creator = c
			break
		}
	}
	return creator, nil
}

func (r *mutationResolver) InsertVideo(ctx context.Context, input model.NewVideo) (*models.Video, error) {
	video := models.Video{
		Title:          input.Title,
		Description:    input.Description,
		Duration:       input.Duration,
		DateUpload:     input.DateUpload,
		DatePublish:    input.DatePublish,
		PhotoThumbnail: input.PhotoThumbnail,
		VideoFile:      input.VideoFile,
		Restricted:     input.Restricted,
		Privacy:        input.Privacy,
		Premium:        input.Premium,
		Status:         input.Status,
		Viewer:         input.Viewer,
		CategoryID:     input.CategoryID,
		CreatorID:      input.CreatorID,
		Location:       input.Location,
	}

	_, err := r.DB.Model(&video).Insert()

	if err != nil {
		return nil, errors.New("insert new creator failed")
	}

	return &video, nil
}

func (r *mutationResolver) InsertCreator(ctx context.Context, input model.NewCreator) (*models.Creator, error) {
	creator := models.Creator{
		Name:             input.Name,
		Email:            input.Email,
		JoinDate:         input.JoinDate,
		Description:      input.Description,
		PhotoProfile:     input.PhotoProfile,
		PhotoBackground:  input.PhotoBackground,
		MembershipStatus: input.MembershipStatus,
	}

	_, err := r.DB.Model(&creator).Insert()

	if err != nil {
		return nil, errors.New("insert new creator failed")
	}

	return &creator, nil
}

func (r *mutationResolver) InsertComment(ctx context.Context, input model.NewComment) (*models.Comment, error) {
	comment := models.Comment{
		Content:    input.Content,
		DateUpload: input.DateUpload,
		CreatorID:  input.CreatorID,
		VideoID:    input.VideoID,
	}

	_, err := r.DB.Model(&comment).Insert()

	if err != nil {
		return nil, errors.New("Insert new comment failed")
	}

	return &comment, nil
}

func (r *mutationResolver) InsertReply(ctx context.Context, input model.NewReply) (*models.Reply, error) {
	reply := models.Reply{
		Content:    input.Content,
		DateUpload: input.DateUpload,
		CreatorID:  input.CreatorID,
		CommentID:  input.CommentID,
	}

	_, err := r.DB.Model(&reply).Insert()

	if err != nil {
		return nil, errors.New("Insert new reply failed")
	}

	return &reply, nil
}

func (r *mutationResolver) InsertSubscribe(ctx context.Context, input model.NewSubscriber) (*models.Subscriber, error) {
	subscriber := models.Subscriber{
		ChannelID: input.ChannelID,
		CreatorID: input.CreatorID,
	}

	_, err := r.DB.Model(&subscriber).Insert()

	if err != nil {
		return nil, errors.New("Insert new Subscriber failed")
	}

	return &subscriber, nil
}

func (r *mutationResolver) DeleteSubscribe(ctx context.Context, creatorID int, channelID int) (bool, error) {
	var subscriber models.Subscriber

	err := r.DB.Model(&subscriber).Where("creator_id = ?", creatorID).Where("channel_id = ?", channelID).First()

	if err != nil {
		return false, errors.New("Subscriber not found")
	}

	_, deleteError := r.DB.Model(&subscriber).Where("creator_id = ?", creatorID).Where("channel_id = ?", channelID).Delete()

	if deleteError != nil {
		return false, errors.New("delete subscriber failed")
	}
	return true, nil
}

func (r *mutationResolver) InsertMembership(ctx context.Context, input model.NewMembership) (*models.Membership, error) {
	membership := models.Membership{
		CreatorID:   input.CreatorID,
		Type:        input.Type,
		Date:        input.Date,
		ExpiredDate: input.ExpiredDate,
		Status:      input.Status,
	}

	_, err := r.DB.Model(&membership).Insert()

	if err != nil {
		return nil, errors.New("Insert new membership failed")
	}

	return &membership, nil
}

func (r *mutationResolver) UpdateMembershipCreator(ctx context.Context, creatorID int, input model.UpdateMembershipStatusCreator) (*models.Creator, error) {
	var creator models.Creator

	err := r.DB.Model(&creator).Where("id = ?", creatorID).First()

	if err != nil {
		return nil, errors.New("Creator not found")
	}

	creator.MembershipStatus = input.MembershipStatus

	_, updateErr := r.DB.Model(&creator).Where("id = ?", creatorID).Update()

	if updateErr != nil {
		return nil, errors.New("Update Creator Membership Status failed")
	}

	return &creator, nil
}

func (r *mutationResolver) InsertLikeCommunity(ctx context.Context, input model.NewLikeCommunity) (*models.Feedbackcommunity, error) {
	fbCommunity := models.Feedbackcommunity{
		CommunityID: input.CommunityID,
		LikerID:     input.LikerID,
		Status:      input.Status,
	}

	_, err := r.DB.Model(&fbCommunity).Insert()

	if err != nil {
		return nil, errors.New("Insert new fbCommunity failed")
	}

	return &fbCommunity, nil
}

func (r *mutationResolver) UpdateLikeCommunity(ctx context.Context, communityID int, likerID int, input model.UpdateLikeCommunity) (*models.Feedbackcommunity, error) {
	var fbcom models.Feedbackcommunity

	err := r.DB.Model(&fbcom).Where("community_id = ?", communityID).Where("liker_id = ?", likerID).First()

	if err != nil {
		return nil, errors.New("FeedbackCommunity not found")
	}

	fbcom.Status = input.Status

	_, updateErr := r.DB.Model(&fbcom).Where("community_id = ?", communityID).Where("liker_id = ?", likerID).Update()

	if updateErr != nil {
		return nil, errors.New("Update FeedbackCommunity Status failed")
	}

	return &fbcom, nil
}

func (r *mutationResolver) DeleteLikeCommunity(ctx context.Context, communityID int, likerID int) (bool, error) {
	var fbcom models.Feedbackcommunity

	err := r.DB.Model(&fbcom).Where("community_id = ?", communityID).Where("liker_id = ?", likerID).First()

	if err != nil {
		return false, errors.New("FbCom not found")
	}

	_, deleteError := r.DB.Model(&fbcom).Where("community_id = ?", communityID).Where("liker_id = ?", likerID).Delete()

	if deleteError != nil {
		return false, errors.New("delete fbcom failed")
	}
	return true, nil
}

func (r *mutationResolver) InsertLikeVideo(ctx context.Context, input model.NewLikeVideo) (*models.Feedbackvideo, error) {
	fbVideo := models.Feedbackvideo{
		VideoID: input.VideoID,
		LikerID: input.LikerID,
		Status:  input.Status,
	}

	_, err := r.DB.Model(&fbVideo).Insert()

	if err != nil {
		return nil, errors.New("Insert new fbVideo failed")
	}

	return &fbVideo, nil
}

func (r *mutationResolver) UpdateLikeVideo(ctx context.Context, videoID int, likerID int, input model.UpdateLikeVideo) (*models.Feedbackvideo, error) {
	var fbVideo models.Feedbackvideo

	err := r.DB.Model(&fbVideo).Where("video_id = ?", videoID).Where("liker_id = ?", likerID).First()

	if err != nil {
		return nil, errors.New("fbVideo not found")
	}

	fbVideo.Status = input.Status

	_, updateErr := r.DB.Model(&fbVideo).Where("video_id = ?", videoID).Where("liker_id = ?", likerID).Update()

	if updateErr != nil {
		return nil, errors.New("Update fbVideo Status failed")
	}

	return &fbVideo, nil
}

func (r *mutationResolver) DeleteLikeVideo(ctx context.Context, videoID int, likerID int) (bool, error) {
	var fbVideo models.Feedbackvideo

	err := r.DB.Model(&fbVideo).Where("video_id = ?", videoID).Where("liker_id = ?", likerID).First()

	if err != nil {
		return false, errors.New("fbVideo not found")
	}

	_, deleteError := r.DB.Model(&fbVideo).Where("video_id = ?", videoID).Where("liker_id = ?", likerID).Delete()

	if deleteError != nil {
		return false, errors.New("delete fbVideo failed")
	}
	return true, nil
}

func (r *mutationResolver) InsertLikeComment(ctx context.Context, input model.NewLikeComment) (*models.Feedbackcomment, error) {
	fbComment := models.Feedbackcomment{
		CommentID: input.CommentID,
		LikerID:   input.LikerID,
		Status:    input.Status,
	}

	_, err := r.DB.Model(&fbComment).Insert()

	if err != nil {
		return nil, errors.New("Insert new fbComment failed")
	}

	return &fbComment, nil
}

func (r *mutationResolver) UpdateLikeComment(ctx context.Context, commentID int, likerID int, input model.UpdateLikeComment) (*models.Feedbackcomment, error) {
	var fbComment models.Feedbackcomment

	err := r.DB.Model(&fbComment).Where("comment_id = ?", commentID).Where("liker_id = ?", likerID).First()

	if err != nil {
		return nil, errors.New("fbComment not found")
	}

	fbComment.Status = input.Status

	_, updateErr := r.DB.Model(&fbComment).Where("comment_id = ?", commentID).Where("liker_id = ?", likerID).Update()

	if updateErr != nil {
		return nil, errors.New("Update fbComment Status failed")
	}

	return &fbComment, nil
}

func (r *mutationResolver) DeleteLikeComment(ctx context.Context, commentID int, likerID int) (bool, error) {
	var fbComment models.Feedbackcomment

	err := r.DB.Model(&fbComment).Where("comment_id = ?", commentID).Where("liker_id = ?", likerID).First()

	if err != nil {
		return false, errors.New("fbComment not found")
	}

	_, deleteError := r.DB.Model(&fbComment).Where("comment_id = ?", commentID).Where("liker_id = ?", likerID).Delete()

	if deleteError != nil {
		return false, errors.New("delete fbComment failed")
	}
	return true, nil
}

func (r *mutationResolver) InsertUserPlaylist(ctx context.Context, input model.NewUserPlaylist) (*model.Userplaylist, error) {
	userplaylist := model.Userplaylist{
		CreatorID:  input.CreatorID,
		PlaylistID: input.PlaylistID,
	}

	_, err := r.DB.Model(&userplaylist).Insert()

	if err != nil {
		return nil, errors.New("Insert new userplaylist failed")
	}

	return &userplaylist, nil
}

func (r *mutationResolver) DeleteUserPlaylist(ctx context.Context, creatorID int, playlistID int) (bool, error) {
	var userplaylist model.Userplaylist

	_, deleteError := r.DB.Model(&userplaylist).Where("creator_id = ?", creatorID).Where("playlist_id = ?", playlistID).Delete()

	if deleteError != nil {
		return false, errors.New("delete userplaylist failed")
	}
	return true, nil
}

func (r *mutationResolver) DeleteAllVideoPlaylist(ctx context.Context, playlistID int) (bool, error) {
	var playlist models.Playlistdetail

	_, deleteError := r.DB.Model(&playlist).Where("playlist_id = ?", playlistID).Delete()

	if deleteError != nil {
		return false, errors.New("delete playlistdetail failed")
	}
	return true, nil
}

func (r *mutationResolver) UpdateEditPlaylist(ctx context.Context, playlistID int, input model.UpdateDate) (*models.Playlist, error) {
	var playlist models.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", playlistID).First()

	if err != nil {
		return nil, errors.New("playlist not found")
	}

	playlist.Date = input.Date

	_, updateErr := r.DB.Model(&playlist).Where("id = ?", playlistID).Update()

	if updateErr != nil {
		return nil, errors.New("Update playlist Status failed")
	}

	return &playlist, nil
}

func (r *mutationResolver) DeleteVideoPlaylist(ctx context.Context, playlistID int, videoID int) (bool, error) {
	var playlist models.Playlistdetail

	_, deleteError := r.DB.Model(&playlist).Where("playlist_id = ?", playlistID).Where("video_id = ?", videoID).Delete()

	if deleteError != nil {
		return false, errors.New("delete playlistdetail failed")
	}
	return true, nil
}

func (r *mutationResolver) InsertDetailPlaylist(ctx context.Context, input model.NewDetailPlaylist) (*models.Playlistdetail, error) {
	playlist := models.Playlistdetail{
		PlaylistID: input.PlaylistID,
		VideoID:    input.VideoID,
	}

	_, err := r.DB.Model(&playlist).Insert()

	if err != nil {
		return nil, errors.New("Insert new playlistdetail failed")
	}

	return &playlist, nil
}

func (r *mutationResolver) InsertNewPlaylist(ctx context.Context, input model.NewPlaylist) (*models.Playlist, error) {
	playlist := models.Playlist{
		Title:       input.Title,
		Date:        input.Date,
		Description: input.Description,
		Viewer:      input.Viewer,
		CreatorID:   input.CreatorID,
		Type:        input.Type,
	}

	_, err := r.DB.Model(&playlist).Insert()

	if err != nil {
		return nil, errors.New("Insert new playlist failed")
	}

	return &playlist, nil
}

func (r *mutationResolver) DeletePlaylistByPlaylistID(ctx context.Context, playlistID int) (bool, error) {
	var playlist models.Playlist

	_, deleteError := r.DB.Model(&playlist).Where("id = ?", playlistID).Delete()

	if deleteError != nil {
		return false, errors.New("delete playlist failed")
	}
	return true, nil
}

func (r *mutationResolver) UpdateBanner(ctx context.Context, creatorID int, input model.UpdateBanner) (*models.Creator, error) {
	var creator models.Creator

	err := r.DB.Model(&creator).Where("id = ?", creatorID).First()

	if err != nil {
		return nil, errors.New("creator not found")
	}

	creator.PhotoBackground = input.PhotoBackground

	_, updateErr := r.DB.Model(&creator).Where("id = ?", creatorID).Update()

	if updateErr != nil {
		return nil, errors.New("Update creator Status failed")
	}

	return &creator, nil
}

func (r *mutationResolver) UpdateIcon(ctx context.Context, creatorID int, input model.UpdateIcon) (*models.Creator, error) {
	var creator models.Creator

	err := r.DB.Model(&creator).Where("id = ?", creatorID).First()

	if err != nil {
		return nil, errors.New("creator not found")
	}

	creator.PhotoProfile = input.PhotoProfile

	_, updateErr := r.DB.Model(&creator).Where("id = ?", creatorID).Update()

	if updateErr != nil {
		return nil, errors.New("Update creator Status failed")
	}

	return &creator, nil
}

func (r *mutationResolver) InsertCommunity(ctx context.Context, input model.NewCommunity) (*models.Community, error) {
	community := models.Community{
		Title:       input.Title,
		Description: input.Description,
		PostImage:   input.PostImage,
		Date:        input.Date,
		CreatorID:   input.CreatorID,
	}

	_, err := r.DB.Model(&community).Insert()

	if err != nil {
		return nil, errors.New("Insert new community failed")
	}

	return &community, nil
}

func (r *mutationResolver) UpdateTitlePlaylist(ctx context.Context, playlistID int, input model.UpdateTitlePlaylist) (*models.Playlist, error) {
	var playlist models.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", playlistID).First()

	if err != nil {
		return nil, errors.New("playlist not found")
	}

	playlist.Title = input.Title

	_, updateErr := r.DB.Model(&playlist).Where("id = ?", playlistID).Update()

	if updateErr != nil {
		return nil, errors.New("Update playlist title failed")
	}

	return &playlist, nil
}

func (r *mutationResolver) UpdateDescriptionPlaylist(ctx context.Context, playlistID int, input model.UpdateDescriptionPlaylist) (*models.Playlist, error) {
	var playlist models.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", playlistID).First()

	if err != nil {
		return nil, errors.New("playlist not found")
	}

	playlist.Description = input.Description

	_, updateErr := r.DB.Model(&playlist).Where("id = ?", playlistID).Update()

	if updateErr != nil {
		return nil, errors.New("Update playlist description failed")
	}

	return &playlist, nil
}

func (r *mutationResolver) UpdateTypePlaylist(ctx context.Context, playlistID int, input model.UpdateTypePlaylist) (*models.Playlist, error) {
	var playlist models.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", playlistID).First()

	if err != nil {
		return nil, errors.New("playlist not found")
	}

	playlist.Type = input.Type

	_, updateErr := r.DB.Model(&playlist).Where("id = ?", playlistID).Update()

	if updateErr != nil {
		return nil, errors.New("Update playlist type failed")
	}

	return &playlist, nil
}

func (r *mutationResolver) DeleteVideo(ctx context.Context, videoID int) (bool, error) {
	var video models.Video

	_, deleteError := r.DB.Model(&video).Where("id = ?", videoID).Delete()

	if deleteError != nil {
		return false, errors.New("delete video failed")
	}
	return true, nil
}

func (r *mutationResolver) UpdateVideo(ctx context.Context, id int, input model.UpdateVideo) (*models.Video, error) {
	var video models.Video

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("video not found")
	}

	video.Title = input.Title
	video.Description = input.Description
	video.PhotoThumbnail = input.PhotoThumbnail
	video.Privacy = input.Privacy

	_, updateErr := r.DB.Model(&video).Where("id = ?", id).Update()

	if updateErr != nil {
		return nil, errors.New("Update video type failed")
	}

	return &video, nil
}

func (r *notifpostResolver) Community(ctx context.Context, obj *models.Notifpost) (*models.Community, error) {
	var communities []*models.Community
	var community *models.Community
	err := r.DB.Model(&communities).Select()
	if err != nil {
		return nil, errors.New("failed to select community :(")
	}
	for _, c := range communities {
		if c.ID == obj.CommunityID {
			community = c
			break
		}
	}
	return community, nil
}

func (r *notifvideoResolver) Video(ctx context.Context, obj *models.Notifvideo) (*models.Video, error) {
	var videos []*models.Video
	var video *models.Video
	err := r.DB.Model(&videos).Select()
	if err != nil {
		return nil, errors.New("failed to select video :(")
	}
	for _, v := range videos {
		if v.ID == obj.VideoID {
			video = v
			break
		}
	}
	return video, nil
}

func (r *playlistResolver) Creator(ctx context.Context, obj *models.Playlist) (*models.Creator, error) {
	var creators []*models.Creator
	var creator *models.Creator
	err := r.DB.Model(&creators).Select()
	if err != nil {
		return nil, errors.New("failed to select creator :(")
	}
	for _, c := range creators {
		if c.ID == obj.CreatorID {
			creator = c
			break
		}
	}
	return creator, nil
}

func (r *playlistdetailResolver) Playlist(ctx context.Context, obj *models.Playlistdetail) (*models.Playlist, error) {
	var playlists []*models.Playlist
	var playlist *models.Playlist
	err := r.DB.Model(&playlists).Select()
	if err != nil {
		return nil, errors.New("failed to select creator :(")
	}
	for _, c := range playlists {
		if c.ID == obj.PlaylistID {
			playlist = c
			break
		}
	}
	return playlist, nil
}

func (r *playlistdetailResolver) Video(ctx context.Context, obj *models.Playlistdetail) (*models.Video, error) {
	var videos []*models.Video
	var video *models.Video
	err := r.DB.Model(&videos).Select()
	if err != nil {
		return nil, errors.New("failed to select video :(")
	}
	for _, v := range videos {
		if v.ID == obj.VideoID {
			video = v
			break
		}
	}
	return video, nil
}

func (r *queryResolver) GetNewIDVideoUpload(ctx context.Context) ([]*models.Video, error) {
	var allVideos []*models.Video
	err := r.DB.Model(&allVideos).Order("id DESC").First()
	if err != nil {
		return nil, errors.New("failed to select video :(")
	}
	return allVideos, nil
}

func (r *queryResolver) Videos(ctx context.Context) ([]*models.Video, error) {
	var allVideos []*models.Video
	err := r.DB.Model(&allVideos).Select()
	if err != nil {
		return nil, errors.New("failed to select video :(")
	}
	return allVideos, nil
}

func (r *queryResolver) Creators(ctx context.Context) ([]*models.Creator, error) {
	var allCreator []*models.Creator
	err := r.DB.Model(&allCreator).Select()
	if err != nil {
		return nil, errors.New("failed to select creator :(")
	}
	return allCreator, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*models.Category, error) {
	var categories []*models.Category
	err := r.DB.Model(&categories).Select()
	if err != nil {
		return nil, errors.New("failed to select categories :(")
	}
	return categories, nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]*models.Comment, error) {
	var comments []*models.Comment
	err := r.DB.Model(&comments).Select()
	if err != nil {
		return nil, errors.New("failed to select comments :(")
	}
	return comments, nil
}

func (r *queryResolver) Replies(ctx context.Context) ([]*models.Reply, error) {
	var replies []*models.Reply
	err := r.DB.Model(&replies).Select()
	if err != nil {
		return nil, errors.New("failed to select comments :(")
	}
	return replies, nil
}

func (r *queryResolver) Subscribers(ctx context.Context) ([]*models.Subscriber, error) {
	var subcribers []*models.Subscriber
	err := r.DB.Model(&subcribers).Select()
	if err != nil {
		return nil, errors.New("failed to select subscriber :(")
	}
	return subcribers, nil
}

func (r *queryResolver) Memberships(ctx context.Context) ([]*models.Membership, error) {
	var membership []*models.Membership
	err := r.DB.Model(&membership).Select()
	if err != nil {
		return nil, errors.New("failed to select subscriber :(")
	}
	return membership, nil
}

func (r *queryResolver) Playlists(ctx context.Context) ([]*models.Playlist, error) {
	var playlist []*models.Playlist
	err := r.DB.Model(&playlist).Select()
	if err != nil {
		return nil, errors.New("failed to select playlist :(")
	}
	return playlist, nil
}

func (r *queryResolver) Playlistdetails(ctx context.Context) ([]*models.Playlistdetail, error) {
	var playlistdetail []*models.Playlistdetail
	err := r.DB.Model(&playlistdetail).Select()
	if err != nil {
		return nil, errors.New("failed to select playlistdetail :(")
	}
	return playlistdetail, nil
}

func (r *queryResolver) Communities(ctx context.Context) ([]*models.Community, error) {
	var community []*models.Community
	err := r.DB.Model(&community).Select()
	if err != nil {
		return nil, errors.New("failed to select comunity :(")
	}
	return community, nil
}

func (r *queryResolver) Feedbackcommunities(ctx context.Context) ([]*models.Feedbackcommunity, error) {
	var fbcommunity []*models.Feedbackcommunity
	err := r.DB.Model(&fbcommunity).Select()
	if err != nil {
		return nil, errors.New("failed to select Feedbackcommunity :(")
	}
	return fbcommunity, nil
}

func (r *queryResolver) Feedbackvideos(ctx context.Context) ([]*models.Feedbackvideo, error) {
	var fbVideo []*models.Feedbackvideo
	err := r.DB.Model(&fbVideo).Select()
	if err != nil {
		return nil, errors.New("failed to select fbVideo :(")
	}
	return fbVideo, nil
}

func (r *queryResolver) Feedbackcomments(ctx context.Context) ([]*models.Feedbackcomment, error) {
	var fbComment []*models.Feedbackcomment
	err := r.DB.Model(&fbComment).Select()
	if err != nil {
		return nil, errors.New("failed to select fbComment :(")
	}
	return fbComment, nil
}

func (r *queryResolver) Userplaylists(ctx context.Context) ([]*model.Userplaylist, error) {
	var uplaylist []*model.Userplaylist
	err := r.DB.Model(&uplaylist).Select()
	if err != nil {
		return nil, errors.New("failed to select uplaylist :(")
	}
	return uplaylist, nil
}

func (r *queryResolver) Notifvideos(ctx context.Context) ([]*models.Notifvideo, error) {
	var nv []*models.Notifvideo
	err := r.DB.Model(&nv).Select()
	if err != nil {
		return nil, errors.New("failed to select nv :(")
	}
	return nv, nil
}

func (r *queryResolver) Notifposts(ctx context.Context) ([]*models.Notifpost, error) {
	var np []*models.Notifpost
	err := r.DB.Model(&np).Select()
	if err != nil {
		return nil, errors.New("failed to select np :(")
	}
	return np, nil
}

func (r *queryResolver) GetHomeVideo(ctx context.Context) ([]*models.Video, error) {
	var allVideos []*models.Video
	err := r.DB.Model(&allVideos).Where("privacy = 2").Select()
	if err != nil {
		return nil, errors.New("failed to select video :(")
	}
	return allVideos, nil
}

func (r *queryResolver) GetVideoByID(ctx context.Context, id int) ([]*models.Video, error) {
	var videos []*models.Video
	err := r.DB.Model(&videos).Where("id = ?", id).Select()
	if err != nil {
		return nil, errors.New("failed to select video :(")
	}
	return videos, nil
}

func (r *queryResolver) SearchingVideo(ctx context.Context, search string) ([]*models.Video, error) {
	var allVideos []*models.Video
	var videos []*models.Video
	err := r.DB.Model(&allVideos).Where("privacy = 2").Select()
	if err != nil {
		return nil, errors.New("failed to select video :(")
	}
	for i := range allVideos {
		if strings.Contains(strings.ToLower(allVideos[i].Title), strings.ToLower(search)) {
			videos = append(videos, allVideos[i])
		}
	}
	return videos, nil
}

func (r *queryResolver) SearchingCreator(ctx context.Context, search string) ([]*models.Creator, error) {
	var allCreators []*models.Creator
	var creators []*models.Creator
	err := r.DB.Model(&allCreators).Select()
	if err != nil {
		return nil, errors.New("failed to select video :(")
	}
	for i := range allCreators {
		if strings.Contains(strings.ToLower(allCreators[i].Name), strings.ToLower(search)) {
			creators = append(creators, allCreators[i])
		}
	}
	return creators, nil
}

func (r *queryResolver) SearchingPlaylists(ctx context.Context, search string) ([]*models.Playlist, error) {
	var allPlaylists []*models.Playlist
	var playlists []*models.Playlist
	err := r.DB.Model(&allPlaylists).Where("type = 2").Select()
	if err != nil {
		return nil, errors.New("failed to select playlist :(")
	}
	for i := range allPlaylists {
		if strings.Contains(strings.ToLower(allPlaylists[i].Title), strings.ToLower(search)) {
			playlists = append(playlists, allPlaylists[i])
		}
	}
	return playlists, nil
}

func (r *queryResolver) GetVideo(ctx context.Context, date string) ([]*models.Video, error) {
	var allVideos []*models.Video
	err := r.DB.Model(&allVideos).Where("date_upload = ?", date).Select()
	if err != nil {
		return nil, errors.New("failed to select video :(")
	}
	return allVideos, nil
}

func (r *queryResolver) GetCreator(ctx context.Context, email string) ([]*models.Creator, error) {
	var allCreator []*models.Creator
	err := r.DB.Model(&allCreator).Where("email = ?", email).Select()
	if err != nil {
		return nil, errors.New("failed to select Creator :(")
	}
	return allCreator, nil
}

func (r *queryResolver) GetCreatorByID(ctx context.Context, id int) ([]*models.Creator, error) {
	var allCreator []*models.Creator
	err := r.DB.Model(&allCreator).Where("id = ?", id).Select()
	if err != nil {
		return nil, errors.New("failed to select Creator :(")
	}
	return allCreator, nil
}

func (r *queryResolver) GetTrendingVideos(ctx context.Context, d1 string, d2 string, d3 string, d4 string, d5 string, d6 string, d7 string) ([]*models.Video, error) {
	var allVideos []*models.Video
	err := r.DB.Model(&allVideos).Where("privacy = 2").WhereGroup(func(q *orm.Query) (*orm.Query, error) {
		q = q.WhereOr("date_upload = ?", d1).WhereOr("date_upload = ?", d2).WhereOr("date_upload = ?", d3).WhereOr("date_upload = ?", d4).WhereOr("date_upload = ?", d5).WhereOr("date_upload = ?", d6).WhereOr("date_upload = ?", d7)
		return q, nil
	}).Order("viewer DESC").Limit(20).Select()
	if err != nil {
		return nil, errors.New("failed to select video :(")
	}
	return allVideos, nil
}

func (r *queryResolver) GetVideoByCategoryID(ctx context.Context, categoryID int) ([]*models.Video, error) {
	var allVideos []*models.Video
	err := r.DB.Model(&allVideos).Where("privacy = 2").Where("category_id = ?", categoryID).Order("viewer DESC").Select()
	if err != nil {
		return nil, errors.New("failed to select video :(")
	}
	return allVideos, nil
}

func (r *queryResolver) GetCommentByID(ctx context.Context, id int) ([]*models.Comment, error) {
	var allComments []*models.Comment
	err := r.DB.Model(&allComments).Where("video_id = ?", id).Select()
	if err != nil {
		return nil, errors.New("failed to select comments :(")
	}
	return allComments, nil
}

func (r *queryResolver) GetReplyByID(ctx context.Context, id int) ([]*models.Reply, error) {
	var allReplies []*models.Reply
	err := r.DB.Model(&allReplies).Where("comment_id = ?", id).Select()
	if err != nil {
		return nil, errors.New("failed to select replies :(")
	}
	return allReplies, nil
}

func (r *queryResolver) GetSubscriberByCreatorID(ctx context.Context, creatorID int) ([]*models.Subscriber, error) {
	var subcribers []*models.Subscriber
	err := r.DB.Model(&subcribers).Where("creator_id = ?", creatorID).Select()
	if err != nil {
		return nil, errors.New("failed to select subscriber :(")
	}
	return subcribers, nil
}

func (r *queryResolver) GetMembershipByCreatorID(ctx context.Context, creatorID int) ([]*models.Membership, error) {
	var membership []*models.Membership
	err := r.DB.Model(&membership).Where("creator_id = ?", creatorID).Order("id DESC").Select()
	if err != nil {
		return nil, errors.New("failed to select membership :(")
	}
	return membership, nil
}

func (r *queryResolver) GetPlaylistByID(ctx context.Context, id int) ([]*models.Playlist, error) {
	var playlist []*models.Playlist
	err := r.DB.Model(&playlist).Where("id = ?", id).Select()
	if err != nil {
		return nil, errors.New("failed to select playlist :(")
	}
	return playlist, nil
}

func (r *queryResolver) GetPlaylistdetailsByID(ctx context.Context, id int) ([]*models.Playlistdetail, error) {
	var playlistdetail []*models.Playlistdetail
	err := r.DB.Model(&playlistdetail).Where("playlist_id = ?", id).Select()
	if err != nil {
		return nil, errors.New("failed to select playlist video :(")
	}
	return playlistdetail, nil
}

func (r *queryResolver) GetSubscriberByChannelID(ctx context.Context, channelID int) ([]*models.Subscriber, error) {
	var subcribers []*models.Subscriber
	err := r.DB.Model(&subcribers).Where("channel_id = ?", channelID).Select()
	if err != nil {
		return nil, errors.New("failed to select subscriber :(")
	}
	return subcribers, nil
}

func (r *queryResolver) GetPlaylistByCreatorID(ctx context.Context, creatorID int) ([]*models.Playlist, error) {
	var playlist []*models.Playlist
	err := r.DB.Model(&playlist).Where("creator_id = ?", creatorID).Order("id DESC").Select()
	if err != nil {
		return nil, errors.New("failed to select playlist :(")
	}
	return playlist, nil
}

func (r *queryResolver) GetCommunityByCreatorID(ctx context.Context, creatorID int) ([]*models.Community, error) {
	var community []*models.Community
	err := r.DB.Model(&community).Where("creator_id = ?", creatorID).Order("id DESC").Select()
	if err != nil {
		return nil, errors.New("failed to select community :(")
	}
	return community, nil
}

func (r *queryResolver) GetFeedbackCommunityByCommunityID(ctx context.Context, communityID int) ([]*models.Feedbackcommunity, error) {
	var fbCommunity []*models.Feedbackcommunity
	err := r.DB.Model(&fbCommunity).Where("community_id = ?", communityID).Select()
	if err != nil {
		return nil, errors.New("failed to select Feedbackcommunity :(")
	}
	return fbCommunity, nil
}

func (r *queryResolver) GetFeedbackCommunityByLikerID(ctx context.Context, likerID int) ([]*models.Feedbackcommunity, error) {
	var fbCommunity []*models.Feedbackcommunity
	err := r.DB.Model(&fbCommunity).Where("liker_id = ?", likerID).Select()
	if err != nil {
		return nil, errors.New("failed to select Feedbackcommunity :(")
	}
	return fbCommunity, nil
}

func (r *queryResolver) GetFeedbackVideoByVideoID(ctx context.Context, videoID int) ([]*models.Feedbackvideo, error) {
	var fbVideo []*models.Feedbackvideo
	err := r.DB.Model(&fbVideo).Where("video_id = ?", videoID).Select()
	if err != nil {
		return nil, errors.New("failed to select fbVideo :(")
	}
	return fbVideo, nil
}

func (r *queryResolver) GetFeedbackVideoByLikerID(ctx context.Context, likerID int) ([]*models.Feedbackvideo, error) {
	var fbVideo []*models.Feedbackvideo
	err := r.DB.Model(&fbVideo).Where("liker_id = ?", likerID).Select()
	if err != nil {
		return nil, errors.New("failed to select fbVideo :(")
	}
	return fbVideo, nil
}

func (r *queryResolver) GetFeedbackCommentByCommentID(ctx context.Context, commentID int) ([]*models.Feedbackcomment, error) {
	var fbComment []*models.Feedbackcomment
	err := r.DB.Model(&fbComment).Where("comment_id = ?", commentID).Select()
	if err != nil {
		return nil, errors.New("failed to select fbComment :(")
	}
	return fbComment, nil
}

func (r *queryResolver) GetFeedbackCommentByLikerID(ctx context.Context, likerID int) ([]*models.Feedbackcomment, error) {
	var fbComment []*models.Feedbackcomment
	err := r.DB.Model(&fbComment).Where("liker_id = ?", likerID).Select()
	if err != nil {
		return nil, errors.New("failed to select fbComment :(")
	}
	return fbComment, nil
}

func (r *queryResolver) GetUserPlaylistByCreatorID(ctx context.Context, creatorID int) ([]*model.Userplaylist, error) {
	var uplaylist []*model.Userplaylist
	err := r.DB.Model(&uplaylist).Where("creator_id = ?", creatorID).Select()
	if err != nil {
		return nil, errors.New("failed to select uplaylist :(")
	}
	return uplaylist, nil
}

func (r *queryResolver) GetNotifVideoByCreatorID(ctx context.Context, creatorID int) ([]*models.Notifvideo, error) {
	var nv []*models.Notifvideo
	err := r.DB.Model(&nv).Where("creator_id = ?", creatorID).Order("id DESC").Limit(3).Select()
	if err != nil {
		return nil, errors.New("failed to select nv :(")
	}
	return nv, nil
}

func (r *queryResolver) GetNotifPostByCreatorID(ctx context.Context, creatorID int) ([]*models.Notifpost, error) {
	var np []*models.Notifpost
	err := r.DB.Model(&np).Where("creator_id = ?", creatorID).Order("id DESC").Limit(3).Select()
	if err != nil {
		return nil, errors.New("failed to select np :(")
	}
	return np, nil
}

func (r *replyResolver) Creator(ctx context.Context, obj *models.Reply) (*models.Creator, error) {
	var creators []*models.Creator
	var creator *models.Creator
	err := r.DB.Model(&creators).Select()
	if err != nil {
		return nil, errors.New("failed to select creator :(")
	}
	for _, c := range creators {
		if c.ID == obj.CreatorID {
			creator = c
			break
		}
	}
	return creator, nil
}

func (r *subscriberResolver) Creator(ctx context.Context, obj *models.Subscriber) (*models.Creator, error) {
	var creators []*models.Creator
	var creator *models.Creator
	err := r.DB.Model(&creators).Select()
	if err != nil {
		return nil, errors.New("failed to select creator :(")
	}
	for _, c := range creators {
		if c.ID == obj.ChannelID {
			creator = c
			break
		}
	}
	return creator, nil
}

func (r *videoResolver) Category(ctx context.Context, obj *models.Video) (*models.Category, error) {
	var categories []*models.Category
	var category *models.Category
	err := r.DB.Model(&categories).Select()
	if err != nil {
		return nil, errors.New("failed to select creator :(")
	}
	for _, c := range categories {
		if c.ID == obj.CategoryID {
			category = c
			break
		}
	}
	return category, nil
}

func (r *videoResolver) Creator(ctx context.Context, obj *models.Video) (*models.Creator, error) {
	var creators []*models.Creator
	var creator *models.Creator
	err := r.DB.Model(&creators).Select()
	if err != nil {
		return nil, errors.New("failed to select creator :(")
	}
	for _, c := range creators {
		if c.ID == obj.CreatorID {
			creator = c
			break
		}
	}
	return creator, nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// Community returns generated.CommunityResolver implementation.
func (r *Resolver) Community() generated.CommunityResolver { return &communityResolver{r} }

// Creator returns generated.CreatorResolver implementation.
func (r *Resolver) Creator() generated.CreatorResolver { return &creatorResolver{r} }

// Membership returns generated.MembershipResolver implementation.
func (r *Resolver) Membership() generated.MembershipResolver { return &membershipResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Notifpost returns generated.NotifpostResolver implementation.
func (r *Resolver) Notifpost() generated.NotifpostResolver { return &notifpostResolver{r} }

// Notifvideo returns generated.NotifvideoResolver implementation.
func (r *Resolver) Notifvideo() generated.NotifvideoResolver { return &notifvideoResolver{r} }

// Playlist returns generated.PlaylistResolver implementation.
func (r *Resolver) Playlist() generated.PlaylistResolver { return &playlistResolver{r} }

// Playlistdetail returns generated.PlaylistdetailResolver implementation.
func (r *Resolver) Playlistdetail() generated.PlaylistdetailResolver {
	return &playlistdetailResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Reply returns generated.ReplyResolver implementation.
func (r *Resolver) Reply() generated.ReplyResolver { return &replyResolver{r} }

// Subscriber returns generated.SubscriberResolver implementation.
func (r *Resolver) Subscriber() generated.SubscriberResolver { return &subscriberResolver{r} }

// Video returns generated.VideoResolver implementation.
func (r *Resolver) Video() generated.VideoResolver { return &videoResolver{r} }

type categoryResolver struct{ *Resolver }
type commentResolver struct{ *Resolver }
type communityResolver struct{ *Resolver }
type creatorResolver struct{ *Resolver }
type membershipResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type notifpostResolver struct{ *Resolver }
type notifvideoResolver struct{ *Resolver }
type playlistResolver struct{ *Resolver }
type playlistdetailResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type replyResolver struct{ *Resolver }
type subscriberResolver struct{ *Resolver }
type videoResolver struct{ *Resolver }
