import 'package:flutter/material.dart';
import 'package:camly_app/constants/routing/startup_path.dart';
import 'package:go_router/go_router.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

import '../../constants/routing/auth_path.dart';
import '../../features/auth/presentation/top_page.dart';
import '../../features/home/navigation/home_shell_branch.dart';
import '../../features/home/navigation/routes/home_page_route.dart';
import '../../features/home/navigation/routes/sample_1_page_route.dart';
import '../../features/home/navigation/routes/sample_2_page_route.dart';
import '../../features/setting/navigation/routes/setting_page_route.dart';
import '../../features/setting/navigation/setting_shell_branch.dart';
import '../../features/startup/application/app_start_app_service.dart';
import '../../features/startup/presentation/app_start_app.dart';
import '../widgets/app_navigation_bar.dart';

part 'app_router.g.dart';

@riverpod
GoRouter goRouter(GoRouterRef ref) {
  final appStartupState = ref.watch(appStartupServiceProvider);

  return GoRouter(
    initialLocation: StartupPath.root,
    debugLogDiagnostics: true,
    routes: $appRoutes,

    /// 状態によって遷移先を変更する
    ///
    /// ユーザの状態によって画面を切り替える場合は各ファイルではなく下記redirectで行う。
    /// その他特定のモデルに依存する状態の場合はroute-levelでのredirectを行う。
    redirect: (context, state) {
      if (appStartupState.isLoading || appStartupState.hasError) {
        return StartupPath.root;
      }

      return null;
    },
  );
}

/// アプリの初期化を行うためのルート。
///
/// 処理内容は下記の通り。
/// - ここに処理内容を記述していく。
@TypedGoRoute<StartupRoute>(
  path: StartupPath.root,
)
class StartupRoute extends GoRouteData {
  const StartupRoute();

  @override
  Widget build(BuildContext context, GoRouterState state) =>
      const AppStartupWidget();
}

/// 認証に関連するルート。
@TypedGoRoute<AuthRoute>(
  path: AuthPath.root,
)
class AuthRoute extends GoRouteData {
  const AuthRoute();

  @override
  Widget build(BuildContext context, GoRouterState state) => const TopPage();
}

// メインシェルルート定義
@TypedStatefulShellRoute<MainShellRouteData>(
  branches: <TypedStatefulShellBranch<StatefulShellBranchData>>[
    homeStatefulShellBranch,
    settingStatefulShellBranch,
  ],
)

// メインシェルルートの状態
class MainShellRouteData extends StatefulShellRouteData {
  const MainShellRouteData();

  @override
  Widget builder(
    BuildContext context,
    GoRouterState state,
    StatefulNavigationShell navigationShell,
  ) {
    return AppNavigationBar(
      navigationShell: navigationShell,
    );
  }
}
