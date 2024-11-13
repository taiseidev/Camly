import 'common_path.dart';

/// Startup機能関連で使用される共通パス定数を定義するユーティリティクラス。
///
/// 使用例：
/// ```dart
/// String rootPath = AuthPath.root;
/// ```
/// 詳細については[CommonPath]を参照。
final class AuthPath {
  const AuthPath._();

  static const root = '/auth';
}
