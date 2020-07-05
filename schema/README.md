# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [user.proto](#user.proto)
    - [CreateUserReq](#user.CreateUserReq)
    - [CreateUserReqInfo](#user.CreateUserReqInfo)
    - [CreateUserRes](#user.CreateUserRes)
    - [CurrentUserReq](#user.CurrentUserReq)
    - [GetUserReq](#user.GetUserReq)
    - [LoginReq](#user.LoginReq)
    - [LoginRes](#user.LoginRes)
    - [LogoutReq](#user.LogoutReq)
    - [RefreshIDTokenReq](#user.RefreshIDTokenReq)
    - [RefreshIDTokenRes](#user.RefreshIDTokenRes)
    - [TokenPair](#user.TokenPair)
    - [UpdatePasswordReq](#user.UpdatePasswordReq)
    - [UpdateUserReq](#user.UpdateUserReq)
    - [UpdateUserReqInfo](#user.UpdateUserReqInfo)
    - [User](#user.User)
  
    - [Sex](#user.Sex)
  
    - [UserService](#user.UserService)
  
- [chat.proto](#chat.proto)
    - [CreateMemberReq](#chat.CreateMemberReq)
    - [CreateMessageReq](#chat.CreateMessageReq)
    - [CreateMessageReqInfo](#chat.CreateMessageReqInfo)
    - [CreateRoomReq](#chat.CreateRoomReq)
    - [DeleteMemberReq](#chat.DeleteMemberReq)
    - [GetRoomReq](#chat.GetRoomReq)
    - [IsMemberReq](#chat.IsMemberReq)
    - [ListMembersReq](#chat.ListMembersReq)
    - [ListMembersRes](#chat.ListMembersRes)
    - [ListMessagesReq](#chat.ListMessagesReq)
    - [ListMessagesRes](#chat.ListMessagesRes)
    - [Member](#chat.Member)
    - [Message](#chat.Message)
    - [Room](#chat.Room)
    - [StreamMessageReq](#chat.StreamMessageReq)
  
    - [ChatService](#chat.ChatService)
  
- [event.proto](#event.proto)
    - [ApplyPostCreated](#event.ApplyPostCreated)
    - [ApplyPostDeleted](#event.ApplyPostDeleted)
    - [CreateRoom](#event.CreateRoom)
    - [CreateRoomFailed](#event.CreateRoomFailed)
    - [Event](#event.Event)
    - [PostApproved](#event.PostApproved)
    - [PostDeleted](#event.PostDeleted)
    - [PostRejected](#event.PostRejected)
    - [RoomCreated](#event.RoomCreated)
  
- [post.proto](#post.proto)
    - [ApplyPost](#post.ApplyPost)
    - [BatchGetApplyPostsByPostIDsReq](#post.BatchGetApplyPostsByPostIDsReq)
    - [BatchGetApplyPostsByPostIDsRes](#post.BatchGetApplyPostsByPostIDsRes)
    - [CreateApplyPostReq](#post.CreateApplyPostReq)
    - [CreatePostReq](#post.CreatePostReq)
    - [CreatePostReqInfo](#post.CreatePostReqInfo)
    - [CreatePostRes](#post.CreatePostRes)
    - [DeleteApplyPostReq](#post.DeleteApplyPostReq)
    - [DeletePostReq](#post.DeletePostReq)
    - [DeletePostRes](#post.DeletePostRes)
    - [GetApplyPostReq](#post.GetApplyPostReq)
    - [GetPostReq](#post.GetPostReq)
    - [ListApplyPostsReq](#post.ListApplyPostsReq)
    - [ListApplyPostsReq.Filter](#post.ListApplyPostsReq.Filter)
    - [ListApplyPostsRes](#post.ListApplyPostsRes)
    - [ListPostsReq](#post.ListPostsReq)
    - [ListPostsReq.Filter](#post.ListPostsReq.Filter)
    - [ListPostsRes](#post.ListPostsRes)
    - [Post](#post.Post)
    - [UpdatePostReq](#post.UpdatePostReq)
    - [UpdatePostReqInfo](#post.UpdatePostReqInfo)
  
    - [ListPostsReq.Filter.OrderBy](#post.ListPostsReq.Filter.OrderBy)
    - [ListPostsReq.Filter.SortBy](#post.ListPostsReq.Filter.SortBy)
  
    - [PostService](#post.PostService)
  
- [event.proto](#event.proto)
    - [ApplyPostCreated](#event.ApplyPostCreated)
    - [ApplyPostDeleted](#event.ApplyPostDeleted)
    - [CreateRoom](#event.CreateRoom)
    - [CreateRoomFailed](#event.CreateRoomFailed)
    - [Event](#event.Event)
    - [PostApproved](#event.PostApproved)
    - [PostDeleted](#event.PostDeleted)
    - [PostRejected](#event.PostRejected)
    - [RoomCreated](#event.RoomCreated)
  
- [image.proto](#image.proto)
    - [BatchCreateImagesReq](#image.BatchCreateImagesReq)
    - [BatchCreateImagesRes](#image.BatchCreateImagesRes)
    - [BatchDeleteImagesByOwnerIDsReq](#image.BatchDeleteImagesByOwnerIDsReq)
    - [BatchDeleteImagesReq](#image.BatchDeleteImagesReq)
    - [DeleteImagesByOwnerIDReq](#image.DeleteImagesByOwnerIDReq)
    - [Image](#image.Image)
    - [ImageInfo](#image.ImageInfo)
    - [ListImagesByOwnerIDReq](#image.ListImagesByOwnerIDReq)
    - [ListImagesByOwnerIDRes](#image.ListImagesByOwnerIDRes)
  
    - [OwnerType](#image.OwnerType)
  
    - [ImageService](#image.ImageService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="user.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## user.proto



<a name="user.CreateUserReq"></a>

### CreateUserReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [CreateUserReqInfo](#user.CreateUserReqInfo) |  |  |
| image_chunk | [bytes](#bytes) |  | ~ 64KiB; |






<a name="user.CreateUserReqInfo"></a>

### CreateUserReqInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| password | [string](#string) |  | 6文字以上72文字以下の英数字 |
| name | [string](#string) |  |  |
| introduction | [string](#string) |  |  |
| sex | [Sex](#user.Sex) |  |  |






<a name="user.CreateUserRes"></a>

### CreateUserRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#user.User) |  |  |
| token_pair | [TokenPair](#user.TokenPair) |  |  |






<a name="user.CurrentUserReq"></a>

### CurrentUserReq







<a name="user.GetUserReq"></a>

### GetUserReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="user.LoginReq"></a>

### LoginReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| password | [string](#string) |  | 6文字以上72文字以下の英数字 |






<a name="user.LoginRes"></a>

### LoginRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#user.User) |  |  |
| token_pair | [TokenPair](#user.TokenPair) |  |  |






<a name="user.LogoutReq"></a>

### LogoutReq







<a name="user.RefreshIDTokenReq"></a>

### RefreshIDTokenReq







<a name="user.RefreshIDTokenRes"></a>

### RefreshIDTokenRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token_pair | [TokenPair](#user.TokenPair) |  |  |






<a name="user.TokenPair"></a>

### TokenPair



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id_token | [string](#string) |  |  |
| refresh_token | [string](#string) |  |  |






<a name="user.UpdatePasswordReq"></a>

### UpdatePasswordReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| old_password | [string](#string) |  | 6文字以上72文字以下の英数字 |
| new_password | [string](#string) |  | 6文字以上72文字以下の英数字 |






<a name="user.UpdateUserReq"></a>

### UpdateUserReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [UpdateUserReqInfo](#user.UpdateUserReqInfo) |  |  |
| image_chunk | [bytes](#bytes) |  | ~ 64KiB; |






<a name="user.UpdateUserReqInfo"></a>

### UpdateUserReqInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| name | [string](#string) |  |  |
| introduction | [string](#string) |  |  |






<a name="user.User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| email | [string](#string) |  |  |
| name | [string](#string) |  |  |
| introduction | [string](#string) |  |  |
| sex | [Sex](#user.Sex) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 


<a name="user.Sex"></a>

### Sex


| Name | Number | Description |
| ---- | ------ | ----------- |
| SEX_UNSPECIFIED | 0 |  |
| MALE | 1 |  |
| FEMALE | 2 |  |


 

 


<a name="user.UserService"></a>

### UserService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CurrentUser | [CurrentUserReq](#user.CurrentUserReq) | [User](#user.User) | require id_token |
| GetUser | [GetUserReq](#user.GetUserReq) | [User](#user.User) | no email field |
| CreateUser | [CreateUserReq](#user.CreateUserReq) stream | [CreateUserRes](#user.CreateUserRes) |  |
| UpdateUser | [UpdateUserReq](#user.UpdateUserReq) stream | [User](#user.User) | require id_token |
| UpdatePassword | [UpdatePasswordReq](#user.UpdatePasswordReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | require id_token |
| Login | [LoginReq](#user.LoginReq) | [LoginRes](#user.LoginRes) |  |
| Logout | [LogoutReq](#user.LogoutReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | require refresh_token |
| RefreshIDToken | [RefreshIDTokenReq](#user.RefreshIDTokenReq) | [RefreshIDTokenRes](#user.RefreshIDTokenRes) | require refresh_token |

 



<a name="chat.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chat.proto



<a name="chat.CreateMemberReq"></a>

### CreateMemberReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |






<a name="chat.CreateMessageReq"></a>

### CreateMessageReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [CreateMessageReqInfo](#chat.CreateMessageReqInfo) |  |  |
| image_chunk | [bytes](#bytes) |  | ~ 64KiB; |






<a name="chat.CreateMessageReqInfo"></a>

### CreateMessageReqInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| body | [string](#string) |  |  |
| room_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |






<a name="chat.CreateRoomReq"></a>

### CreateRoomReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| post_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |






<a name="chat.DeleteMemberReq"></a>

### DeleteMemberReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |






<a name="chat.GetRoomReq"></a>

### GetRoomReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [int64](#int64) |  |  |
| post_id | [int64](#int64) |  |  |






<a name="chat.IsMemberReq"></a>

### IsMemberReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [int64](#int64) |  |  |
| post_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |






<a name="chat.ListMembersReq"></a>

### ListMembersReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [int64](#int64) |  |  |






<a name="chat.ListMembersRes"></a>

### ListMembersRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| members | [Member](#chat.Member) | repeated |  |






<a name="chat.ListMessagesReq"></a>

### ListMessagesReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [int64](#int64) |  |  |






<a name="chat.ListMessagesRes"></a>

### ListMessagesRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| messages | [Message](#chat.Message) | repeated |  |






<a name="chat.Member"></a>

### Member



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| room_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="chat.Message"></a>

### Message



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| body | [string](#string) |  |  |
| room_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="chat.Room"></a>

### Room



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| post_id | [int64](#int64) |  |  |
| members | [Member](#chat.Member) | repeated |  |
| messages | [Message](#chat.Message) | repeated |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="chat.StreamMessageReq"></a>

### StreamMessageReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [int64](#int64) |  |  |





 

 

 


<a name="chat.ChatService"></a>

### ChatService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetRoom | [GetRoomReq](#chat.GetRoomReq) | [Room](#chat.Room) |  |
| CreateRoom | [CreateRoomReq](#chat.CreateRoomReq) | [Room](#chat.Room) | user_idでcreateMemberもする |
| IsMember | [IsMemberReq](#chat.IsMemberReq) | [.google.protobuf.BoolValue](#google.protobuf.BoolValue) | ユーザーがルームメンバーか確認。post_idからroomを取る |
| ListMembers | [ListMembersReq](#chat.ListMembersReq) | [ListMembersRes](#chat.ListMembersRes) |  |
| CreateMember | [CreateMemberReq](#chat.CreateMemberReq) | [Member](#chat.Member) |  |
| DeleteMember | [DeleteMemberReq](#chat.DeleteMemberReq) | [.google.protobuf.Empty](#google.protobuf.Empty) |  |
| ListMessages | [ListMessagesReq](#chat.ListMessagesReq) | [ListMessagesRes](#chat.ListMessagesRes) |  |
| CreateMessage | [CreateMessageReq](#chat.CreateMessageReq) stream | [Message](#chat.Message) |  |
| StreamMessage | [StreamMessageReq](#chat.StreamMessageReq) | [Message](#chat.Message) stream |  |

 



<a name="event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## event.proto



<a name="event.ApplyPostCreated"></a>

### ApplyPostCreated



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apply_post | [post.ApplyPost](#post.ApplyPost) |  |  |






<a name="event.ApplyPostDeleted"></a>

### ApplyPostDeleted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apply_post | [post.ApplyPost](#post.ApplyPost) |  |  |






<a name="event.CreateRoom"></a>

### CreateRoom



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| saga_id | [string](#string) |  |  |
| post_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |






<a name="event.CreateRoomFailed"></a>

### CreateRoomFailed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| saga_id | [string](#string) |  |  |
| message | [string](#string) |  |  |






<a name="event.Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| event_type | [string](#string) |  |  |
| aggregate_id | [string](#string) |  |  |
| aggregate_type | [string](#string) |  |  |
| event_data | [bytes](#bytes) |  |  |
| channel | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="event.PostApproved"></a>

### PostApproved



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| saga_id | [string](#string) |  |  |
| post | [post.Post](#post.Post) |  |  |






<a name="event.PostDeleted"></a>

### PostDeleted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| post | [post.Post](#post.Post) |  |  |






<a name="event.PostRejected"></a>

### PostRejected



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| saga_id | [string](#string) |  |  |
| post | [post.Post](#post.Post) |  |  |
| error_message | [string](#string) |  |  |






<a name="event.RoomCreated"></a>

### RoomCreated



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| saga_id | [string](#string) |  |  |
| room | [chat.Room](#chat.Room) |  |  |





 

 

 

 



<a name="post.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## post.proto



<a name="post.ApplyPost"></a>

### ApplyPost



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| post_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="post.BatchGetApplyPostsByPostIDsReq"></a>

### BatchGetApplyPostsByPostIDsReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| post_ids | [int64](#int64) | repeated |  |






<a name="post.BatchGetApplyPostsByPostIDsRes"></a>

### BatchGetApplyPostsByPostIDsRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apply_posts | [ApplyPost](#post.ApplyPost) | repeated |  |






<a name="post.CreateApplyPostReq"></a>

### CreateApplyPostReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| post_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |






<a name="post.CreatePostReq"></a>

### CreatePostReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [CreatePostReqInfo](#post.CreatePostReqInfo) |  |  |
| next_image_signal | [bool](#bool) |  |  |
| image_chunk | [bytes](#bytes) |  | ~ 64KiB |






<a name="post.CreatePostReqInfo"></a>

### CreatePostReqInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  | 20文字以下 |
| content | [string](#string) |  | 2000文字以下 |
| fishing_spot_type_id | [int64](#int64) |  | 1~4 |
| fish_type_ids | [int64](#int64) | repeated | 1~95 一個以上 ユニーク |
| prefecture_id | [int64](#int64) |  | 1~47 |
| meeting_place_id | [string](#string) |  | google place ID, いまはサーバー側では叩かず保存して返すだけ。後々API叩く。 |
| meeting_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| max_apply | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |






<a name="post.CreatePostRes"></a>

### CreatePostRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| post | [Post](#post.Post) |  |  |
| saga_id | [string](#string) |  |  |






<a name="post.DeleteApplyPostReq"></a>

### DeleteApplyPostReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="post.DeletePostReq"></a>

### DeletePostReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="post.DeletePostRes"></a>

### DeletePostRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="post.GetApplyPostReq"></a>

### GetApplyPostReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="post.GetPostReq"></a>

### GetPostReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="post.ListApplyPostsReq"></a>

### ListApplyPostsReq
one of 使える


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filter | [ListApplyPostsReq.Filter](#post.ListApplyPostsReq.Filter) |  |  |






<a name="post.ListApplyPostsReq.Filter"></a>

### ListApplyPostsReq.Filter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  |  |
| post_id | [int64](#int64) |  |  |






<a name="post.ListApplyPostsRes"></a>

### ListApplyPostsRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apply_posts | [ApplyPost](#post.ApplyPost) | repeated |  |






<a name="post.ListPostsReq"></a>

### ListPostsReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filter | [ListPostsReq.Filter](#post.ListPostsReq.Filter) |  |  |
| page_size | [int64](#int64) |  | 30件以下。ゼロ値の場合、デフォルト設定で10件 |
| page_token | [string](#string) |  |  |






<a name="post.ListPostsReq.Filter"></a>

### ListPostsReq.Filter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prefecture_id | [int64](#int64) |  | フロントでenum管理 都道府県コード |
| fishing_spot_type_id | [int64](#int64) |  | フロントでenum管理 1: 陸っぱり, 2: 渓流釣り, 3: 釣り船, 4: 釣り堀 |
| fish_type_ids | [int64](#int64) | repeated | フロントでenum管理 |
| meeting_at_from | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | fromとtoはセットで必要 エラーは出ないけどクエリは無効になる |
| meeting_at_to | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| can_apply | [bool](#bool) |  | trueにすると、応募可能な投稿のみを絞り込める。 |
| order_by | [ListPostsReq.Filter.OrderBy](#post.ListPostsReq.Filter.OrderBy) |  |  |
| sort_by | [ListPostsReq.Filter.SortBy](#post.ListPostsReq.Filter.SortBy) |  |  |
| user_id | [int64](#int64) |  | ここに値が入っているとユーザーの投稿を絞り込める |






<a name="post.ListPostsRes"></a>

### ListPostsRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| posts | [Post](#post.Post) | repeated |  |
| next_page_token | [string](#string) |  |  |






<a name="post.Post"></a>

### Post



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| content | [string](#string) |  |  |
| fishing_spot_type_id | [int64](#int64) |  |  |
| fish_type_ids | [int64](#int64) | repeated |  |
| prefecture_id | [int64](#int64) |  |  |
| meeting_place_id | [string](#string) |  |  |
| meeting_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| user_id | [int64](#int64) |  |  |
| max_apply | [int64](#int64) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="post.UpdatePostReq"></a>

### UpdatePostReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [UpdatePostReqInfo](#post.UpdatePostReqInfo) |  |  |
| next_image_signal | [bool](#bool) |  |  |
| image_chunk | [bytes](#bytes) |  |  |






<a name="post.UpdatePostReqInfo"></a>

### UpdatePostReqInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| content | [string](#string) |  |  |
| fishing_spot_type_id | [int64](#int64) |  |  |
| fish_type_ids | [int64](#int64) | repeated |  |
| prefecture_id | [int64](#int64) |  |  |
| meeting_place_id | [string](#string) |  |  |
| meeting_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| max_apply | [int64](#int64) |  |  |
| image_ids_to_delete | [int64](#int64) | repeated |  |





 


<a name="post.ListPostsReq.Filter.OrderBy"></a>

### ListPostsReq.Filter.OrderBy


| Name | Number | Description |
| ---- | ------ | ----------- |
| DESC | 0 | デフォルト |
| ASC | 1 |  |



<a name="post.ListPostsReq.Filter.SortBy"></a>

### ListPostsReq.Filter.SortBy


| Name | Number | Description |
| ---- | ------ | ----------- |
| CREATED_AT | 0 | デフォルト |
| MEETING_AT | 1 |  |


 

 


<a name="post.PostService"></a>

### PostService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPost | [GetPostReq](#post.GetPostReq) | [Post](#post.Post) | 投稿の詳細を取得 |
| ListPosts | [ListPostsReq](#post.ListPostsReq) | [ListPostsRes](#post.ListPostsRes) | 投稿の絞り込み検索 &amp; ページネーション |
| CreatePost | [CreatePostReq](#post.CreatePostReq) stream | [CreatePostRes](#post.CreatePostRes) | 投稿を作成 |
| UpdatePost | [UpdatePostReq](#post.UpdatePostReq) stream | [Post](#post.Post) | 投稿を更新 |
| DeletePost | [DeletePostReq](#post.DeletePostReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 投稿を削除 |
| GetApplyPost | [GetApplyPostReq](#post.GetApplyPostReq) | [ApplyPost](#post.ApplyPost) |  |
| ListApplyPosts | [ListApplyPostsReq](#post.ListApplyPostsReq) | [ListApplyPostsRes](#post.ListApplyPostsRes) | その投稿の応募情報のリストを取得, そのユーザーの応募した投稿リストを取得 |
| BatchGetApplyPostsByPostIDs | [BatchGetApplyPostsByPostIDsReq](#post.BatchGetApplyPostsByPostIDsReq) | [BatchGetApplyPostsByPostIDsRes](#post.BatchGetApplyPostsByPostIDsRes) |  |
| CreateApplyPost | [CreateApplyPostReq](#post.CreateApplyPostReq) | [ApplyPost](#post.ApplyPost) | ユーザーが投稿に応募する |
| DeleteApplyPost | [DeleteApplyPostReq](#post.DeleteApplyPostReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | ユーザーがその投稿の応募を取り消す |

 



<a name="event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## event.proto



<a name="event.ApplyPostCreated"></a>

### ApplyPostCreated



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apply_post | [post.ApplyPost](#post.ApplyPost) |  |  |






<a name="event.ApplyPostDeleted"></a>

### ApplyPostDeleted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apply_post | [post.ApplyPost](#post.ApplyPost) |  |  |






<a name="event.CreateRoom"></a>

### CreateRoom



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| saga_id | [string](#string) |  |  |
| post_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |






<a name="event.CreateRoomFailed"></a>

### CreateRoomFailed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| saga_id | [string](#string) |  |  |
| message | [string](#string) |  |  |






<a name="event.Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| event_type | [string](#string) |  |  |
| aggregate_id | [string](#string) |  |  |
| aggregate_type | [string](#string) |  |  |
| event_data | [bytes](#bytes) |  |  |
| channel | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="event.PostApproved"></a>

### PostApproved



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| saga_id | [string](#string) |  |  |
| post | [post.Post](#post.Post) |  |  |






<a name="event.PostDeleted"></a>

### PostDeleted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| post | [post.Post](#post.Post) |  |  |






<a name="event.PostRejected"></a>

### PostRejected



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| saga_id | [string](#string) |  |  |
| post | [post.Post](#post.Post) |  |  |
| error_message | [string](#string) |  |  |






<a name="event.RoomCreated"></a>

### RoomCreated



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| saga_id | [string](#string) |  |  |
| room | [chat.Room](#chat.Room) |  |  |





 

 

 

 



<a name="image.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## image.proto



<a name="image.BatchCreateImagesReq"></a>

### BatchCreateImagesReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [ImageInfo](#image.ImageInfo) |  |  |
| chunk | [bytes](#bytes) |  | ~ 64KiB |






<a name="image.BatchCreateImagesRes"></a>

### BatchCreateImagesRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| images | [Image](#image.Image) | repeated |  |






<a name="image.BatchDeleteImagesByOwnerIDsReq"></a>

### BatchDeleteImagesByOwnerIDsReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner_ids | [int64](#int64) | repeated |  |
| owner_type | [OwnerType](#image.OwnerType) |  |  |






<a name="image.BatchDeleteImagesReq"></a>

### BatchDeleteImagesReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [int64](#int64) | repeated |  |






<a name="image.DeleteImagesByOwnerIDReq"></a>

### DeleteImagesByOwnerIDReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner_id | [int64](#int64) |  |  |
| owner_type | [OwnerType](#image.OwnerType) |  |  |






<a name="image.Image"></a>

### Image



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| owner_id | [int64](#int64) |  |  |
| owner_type | [OwnerType](#image.OwnerType) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="image.ImageInfo"></a>

### ImageInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner_id | [int64](#int64) |  |  |
| owner_type | [OwnerType](#image.OwnerType) |  |  |






<a name="image.ListImagesByOwnerIDReq"></a>

### ListImagesByOwnerIDReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner_id | [int64](#int64) |  |  |
| owner_type | [OwnerType](#image.OwnerType) |  |  |






<a name="image.ListImagesByOwnerIDRes"></a>

### ListImagesByOwnerIDRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| images | [Image](#image.Image) | repeated |  |





 


<a name="image.OwnerType"></a>

### OwnerType


| Name | Number | Description |
| ---- | ------ | ----------- |
| OWNERTYPE_UNSPECIFIED | 0 |  |
| POST | 1 |  |
| USER | 2 |  |
| MESSAGE | 3 |  |


 

 


<a name="image.ImageService"></a>

### ImageService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListImagesByOwnerID | [ListImagesByOwnerIDReq](#image.ListImagesByOwnerIDReq) | [ListImagesByOwnerIDRes](#image.ListImagesByOwnerIDRes) |  |
| BatchCreateImages | [BatchCreateImagesReq](#image.BatchCreateImagesReq) stream | [BatchCreateImagesRes](#image.BatchCreateImagesRes) |  |
| BatchDeleteImages | [BatchDeleteImagesReq](#image.BatchDeleteImagesReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 削除するレコードがなくてもエラーなし |
| BatchDeleteImagesByOwnerIDs | [BatchDeleteImagesByOwnerIDsReq](#image.BatchDeleteImagesByOwnerIDsReq) | [.google.protobuf.Empty](#google.protobuf.Empty) | 削除するレコードがなくてもエラーなし |
| DeleteImagesByOwnerID | [DeleteImagesByOwnerIDReq](#image.DeleteImagesByOwnerIDReq) | [.google.protobuf.Empty](#google.protobuf.Empty) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

