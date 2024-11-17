import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../../core/navigation/app_router.dart';
import '../../home/navigation/routes/home_page_route.dart';
import '../application/app_start_app_service.dart';

final class AppStartupWidget extends ConsumerWidget {
  const AppStartupWidget({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final appStartupState = ref.watch(appStartupServiceProvider);

    return appStartupState.when(
      data: (_) => const _AppStartupDataWidget(),
      loading: () => const _AppStartupLoadingWidget(),
      error: (e, _) => _AppStartupErrorWidget(
        message: e.toString(),
        onRetry: () => ref.invalidate(appStartupServiceProvider),
      ),
    );
  }
}

/// データ取得時
final class _AppStartupDataWidget extends StatelessWidget {
  const _AppStartupDataWidget();

  @override
  Widget build(BuildContext context) {
    /// FIXME: AppStartupWidgetのref.listenやaddPostFrameCallbackでgoするとGoRouterが見つからないエラーが発生する
    WidgetsBinding.instance.addPostFrameCallback((_) {
      const HomePageRoute().go(context);
    });

    return const Scaffold(
      body: Center(
        child: CircularProgressIndicator(),
      ),
    );
  }
}

/// ローディング
final class _AppStartupLoadingWidget extends StatelessWidget {
  const _AppStartupLoadingWidget();

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(
        child: CircularProgressIndicator(),
      ),
    );
  }
}

/// エラー
final class _AppStartupErrorWidget extends StatelessWidget {
  _AppStartupErrorWidget({
    required this.message,
    required this.onRetry,
  });

  final String message;
  final VoidCallback onRetry;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text('エラーが発生しました'),
            Text(message),
            ElevatedButton(
              onPressed: onRetry,
              child: const Text('再試行'),
            ),
          ],
        ),
      ),
    );
  }
}
