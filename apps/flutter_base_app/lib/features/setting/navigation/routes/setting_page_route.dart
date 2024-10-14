import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import '../../presentation/pages/setting_page.dart';

class SettingPageRoute extends GoRouteData {
  const SettingPageRoute();

  @override
  Widget build(BuildContext context, GoRouterState state) {
    return const SettingPage();
  }
}
