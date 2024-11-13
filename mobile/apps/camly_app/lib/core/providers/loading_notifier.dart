import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'loading_notifier.g.dart';

@Riverpod(keepAlive: true)
class LoadingNotifier extends _$LoadingNotifier {
  @override
  bool build() => false;

  void show() => state = true;

  void hide() => state = false;
}
