import 'package:camly_app/core/navigation/app_router.dart';
import 'package:camly_app/features/home/navigation/routes/sample_1_page_route.dart';
import 'package:flutter/material.dart';

import '../../navigation/routes/sample_2_page_route.dart';

class HomePage extends StatelessWidget {
  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text('テスト'),
            ElevatedButton(
              onPressed: () => const Sample1PageRoute().go(context),
              child: const Text('Sample1Pageに遷移'),
            ),
            ElevatedButton(
              onPressed: () => const Sample2PageRoute().go(context),
              child: const Text('Sample2Pageに遷移'),
            ),
          ],
        ),
      ),
    );
  }
}
