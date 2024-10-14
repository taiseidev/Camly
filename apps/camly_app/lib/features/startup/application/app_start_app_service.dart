import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'app_start_app_service.g.dart';

// FIXME(onishi): 初期値で保持しておきたい値をここで返却する。仮でbool値を指定
@Riverpod(keepAlive: true)
Future<bool> appStartupService(AppStartupServiceRef ref) async {
  ref.onDispose(() {});

  await Future.delayed(const Duration(seconds: 3));

  return true;
}
