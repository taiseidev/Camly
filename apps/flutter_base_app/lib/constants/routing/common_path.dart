/// アプリケーション全体で使用される共通パス定数を定義するユーティリティクラス。
///
/// 使用例：
/// ```dart
/// String rootPath = CommonPath.root;
/// ```
///
final class CommonPath {
  /// インスタンス化を防ぐためにプライベート化。
  const CommonPath._();

  /// ルートパスは[root]と命名する。
  static const root = '/';

  /// サブルートはメイン機能を示します。
  static const main = 'main';
}
