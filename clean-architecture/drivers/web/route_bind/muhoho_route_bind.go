package route_bind

import (
	"html/template"
	"net/http"

	"github.com/Greek-Academy/software-architecture/clean-architecture/domain/object_values"
	"github.com/Greek-Academy/software-architecture/clean-architecture/drivers/ui"
	"github.com/Greek-Academy/software-architecture/clean-architecture/interface_adapter/controllers"
	"github.com/Greek-Academy/software-architecture/clean-architecture/interface_adapter/presenters"
)

func MuhohoRouteBind(w http.ResponseWriter, r *http.Request) {
	// リクエストメソッドの検証
	if r.Method != http.MethodGet {
		http.Error(w, object_values.MethodNotAllowed().Error(), http.StatusMethodNotAllowed)
		return
	}

	// テンプレートの読み込み
	t, err := template.ParseFiles(ui.MUHOHO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// インターフェースアダプター層へ
	c := controllers.NewMuhohoController(r)
	res := c.Handle()

	// プレゼンターで返却データを生成
	pre := presenters.NewMuhohoPresenter(res)

	// テンプレートへ変数展開
	if err = t.Execute(w, pre); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
