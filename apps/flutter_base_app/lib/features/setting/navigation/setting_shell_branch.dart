// 設定タプの状態
import 'package:go_router/go_router.dart';

import 'routes/setting_page_route.dart';

const settingStatefulShellBranch = TypedStatefulShellBranch<SettingShellBranch>(
  routes: [
    TypedGoRoute<SettingPageRoute>(
      path: '/setting',
    ),
  ],
);

// 設定タブの定義
class SettingShellBranch extends StatefulShellBranchData {
  const SettingShellBranch();
}
