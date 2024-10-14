// ホームタブの定義
import 'package:go_router/go_router.dart';

import 'routes/home_page_route.dart';
import 'routes/sample_1_page_route.dart';
import 'routes/sample_2_page_route.dart';

// ホームタプの状態
const homeStatefulShellBranch = TypedStatefulShellBranch<HomeShellBranch>(
  routes: <TypedRoute<RouteData>>[
    TypedGoRoute<HomePageRoute>(
      path: '/home',
      routes: [
        TypedGoRoute<Sample1PageRoute>(
          path: 'sample1',
        ),
        TypedGoRoute<Sample2PageRoute>(
          path: 'sample2',
        ),
      ],
    ),
  ],
);

class HomeShellBranch extends StatefulShellBranchData {
  const HomeShellBranch();
}
