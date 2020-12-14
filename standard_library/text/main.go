package main

import "fmt"

func main() {
	link := InnerLink(ArticleCommentLinkType)
	fmt.Println(link)
}


type InnerLinkType string

//内链的type
const (
	TagLinkType                InnerLinkType = "tag"                  //标签页   wanxiu://innerlink?type=tag&id={id}
	SeedRecommendLinkType      InnerLinkType = "seedRecommend"        //精选集列表页   wanxiu://innerlink?type=seedRecommend&tagId={tagId}&title={title}
	ArticleNormalLinkType      InnerLinkType = "article_link"         //普通文章   wanxiu://innerlink?type=article&id={id}
	VideoArticleTypeLinkType   InnerLinkType = "video_article"        //视频文章   wanxiu://innerlink?type=video&id={id}&url={url}&tagId={tagId}
	ArticleCommentListLinkType InnerLinkType = "article_comment_list" //文章评论页   wanxiu://innerlink?type=article_comment_list&id={id}
	ArticleCommentLinkType     InnerLinkType = "article_comment"      //文章二级评论页   wanxiu://innerlink?type=article_comment&id={id}&subId={subId}&articleId={articleId}
	RewardLinkType             InnerLinkType = "reward"               //打赏   wanxiu://innerlink?type=reward&id={id}
	RewardListLinkType         InnerLinkType = "rewardlist"           //打赏列表   wanxiu://innerlink?type=rewardlist&id={id}
	CreatPostLinkType          InnerLinkType = "creatpost"            //发帖（弹出窗）   wanxiu://innerlink?type=creatpost
	PostVideoLinkType          InnerLinkType = "bbspost_video"        //视频帖子   wanxiu://innerlink?type=bbspost_video&id={id}&url={url}
	PostNormalLinkType         InnerLinkType = "bbspost_link"         //普通帖子（包含一级评论）   wanxiu://innerlink?type=bbspost_link&id={id}
	// 二级评论页  wanxiu://innerlink?type=bbspost_link&id={id}&comment={comment}

	PostUrlLinkType      InnerLinkType = "bbspost_url"      //旧富文本帖子   wanxiu://innerlink?type=bbspost_url&id={id}
	PostRichTextLinkType InnerLinkType = "bbspost_richtext" //新富文本帖子   wanxiu://innerlink?type=bbspost_richtext&id={id}
	MallHomeLinkType     InnerLinkType = "mall_home"        //商城首页   wanxiu://innerlink?type=mall_home
	MallItemInfoLinkType InnerLinkType = "mall"             //商品详情   wanxiu://innerlink?type=mall&id={id}
	MallOrderLinkType    InnerLinkType = "mall_order"       //商城订单   wanxiu://innerlink?type= mall_order&id={id}
	MatchHomeLinkType    InnerLinkType = "match"            //比赛首页   wanxiu://innerlink?type=mall_home
	MatchLinkLinkType    InnerLinkType = "match_link"       //比赛首页（区分游戏）   wanxiu://innerlink?type=match_link&game={game}
	GameCenterLinkType   InnerLinkType = "gamecenter"       //游戏中心   wanxiu://innerlink?type=gamecenter
	GameInfoLinkType     InnerLinkType = "game"             //游戏详情   wanxiu://innerlink?type=game&id={id}
	CardInfoLinkType     InnerLinkType = "card"             //单卡详情   wanxiu://innerlink?type=card&id={id}&game={game}
	DeckInfoLinkType     InnerLinkType = "deck_detail"      //套牌详情   wanxiu://innerlink?type=deck_detail&id={id}&game={game}
	DeckSetLinkType      InnerLinkType = "deckset"          //套牌集   wanxiu://innerlink?type=deckset&setid={id}&game={game}&name={name}
	OpenPackLinkType     InnerLinkType = "openpack"         //模拟开包   wanxiu://innerlink?type=openpack&game={game}
	ToolLinkType         InnerLinkType = "tool"             //工具栏   wanxiu://innerlink?type=deckset&id={id}&game={game}
	FeedTopTagLinkType   InnerLinkType = "feed_top_tag"     //feedTop标签跳转 wanxiu://innerlink?type=feed_top_tag&id={id}
)

// InnerLink 构建内链
/*func InnerLink(id int, jumpType InnerLinkType, other string) string {
	return fmt.Sprintf("wanxiu://innerlink?type=%s&id=%d&%s", jumpType, id, other)
}*/
//entry为内链中值的键值对，比如 id=123, game=hearthstone
func InnerLink(jumpType InnerLinkType, entry ...string) string {
	format := "wanxiu://innerlink?type=" + string(jumpType)
	for _, e := range entry {
		format += "&" + e
	}
	return format
}
