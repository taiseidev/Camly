import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import '../../presentation/pages/sample_1_page.dart';

class Sample1PageRoute extends GoRouteData {
  const Sample1PageRoute();

  @override
  Widget build(BuildContext context, GoRouterState state) {
    return const Sample1Page();
  }
}
