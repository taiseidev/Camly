import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import '../../presentation/pages/sample_2_page.dart';

class Sample2PageRoute extends GoRouteData {
  const Sample2PageRoute();

  @override
  Widget build(BuildContext context, GoRouterState state) {
    return const Sample2Page();
  }
}
