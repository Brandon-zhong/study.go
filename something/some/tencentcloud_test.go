package some

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tms/v20201229"
	"testing"
)

func TestTencentCloud(t *testing.T) {

	credential := common.NewCredential(
		"AKID3Ik2TsWuiScaq4laWSkyhHoz8CcfDdV9",
		"11dpixV4bBRAPxDViLzqx8ARvbUVDggu",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "tms.tencentcloudapi.com"
	client, _ := tms.NewClient(credential, "ap-beijing", cpf)

	request := tms.NewTextModerationRequest()
	toString := base64.StdEncoding.EncodeToString([]byte("啊~啊~啊~啊~啊~~~啊~啊~啊~啊啊啊啊啊啊，加我好友，给你发优惠券，啊啊啊啊，肢解，大法师的，色情，玩个涩情，尼玛炸了"))

	request.Content = common.StringPtr(toString)

	resp, err := client.TextModeration(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	byte, err := json.Marshal(resp.Response)
	str := string(byte)
	fmt.Println(str)
	fmt.Println("------------")

	var tmp struct {

		// 您在入参时所填入的Biztype参数。 -- 该字段暂未开放。
		BizType *string `json:"BizType" name:"BizType"`

		// 恶意标签，Normal：正常，Porn：色情，Abuse：谩骂，Ad：广告，Custom：自定义词库。
		// 以及其他令人反感、不安全或不适宜的内容类型。
		Label *string `json:"Label" name:"Label"`

		// 建议您拿到判断结果后的执行操作。
		// 建议值，Block：建议屏蔽，Review：建议复审，Pass：建议通过
		Suggestion *string `json:"Suggestion" name:"Suggestion"`

		// 文本命中的关键词信息，用于提示您文本违规的具体原因，可能会返回多个命中的关键词。（如：加我微信）
		// 如返回值为空，Score不为空，即识别结果（Label）是来自于语义模型判断的返回值。
		// 注意：此字段可能返回 null，表示取不到有效值。
		Keywords []*string `json:"Keywords" name:"Keywords"`

		// 机器判断当前分类的置信度，取值范围：0.00~100.00。分数越高，表示越有可能属于当前分类。
		// （如：色情 99.99，则该样本属于色情的置信度非常高。）
		Score *int64 `json:"Score" name:"Score"`

		// 接口识别样本后返回的详细结果。
		// 注意：此字段可能返回 null，表示取不到有效值。
		DetailResults []*tms.DetailResults `json:"DetailResults" name:"DetailResults"`

		// 接口识别样本中存在违规账号风险的检测结果。
		// 注意：此字段可能返回 null，表示取不到有效值。
		RiskDetails []*tms.RiskDetails `json:"RiskDetails" name:"RiskDetails"`

		// 扩展字段，用于特定信息返回，不同客户/Biztype下返回信息不同。
		// 注意：此字段可能返回 null，表示取不到有效值。
		Extra *string `json:"Extra" name:"Extra"`

		// 请求参数中的DataId
		// 注意：此字段可能返回 null，表示取不到有效值。
		DataId *string `json:"DataId" name:"DataId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	}

	err = json.Unmarshal(byte, &tmp)
	if err != nil {
		fmt.Println(err)
	}
	byte, _ = json.Marshal(tmp)
	str = string(byte)
	fmt.Println(str)


	//fmt.Println(util.BytesToString(must.Byte(json.Marshal(resp.Response))))

}
