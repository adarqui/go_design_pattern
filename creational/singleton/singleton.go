package singleton

// 構造体の名前を小文字にすることでパッケージ外へのエクスポートを行わない
// Not exported
type singleton struct {
}

// インスタンスを保持する変数も小文字にすることでエクスポートを行わない
// Variable that holds the instance also not exported
var instance *singleton

// インスタンス取得用の関数のみエクスポートしておき、ここでインスタンスが
// 一意であることを保証する
// Only export GetInstance to acquire the unique singleton instance
func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
