// ナビゲーションバー
import 'dart:ui';

import 'package:flutter/material.dart';
import 'package:flutter_animate/flutter_animate.dart';
import 'package:go_router/go_router.dart';

enum BottomNavigationBarType {
  home,
  search,
  notification,
  setting,
  create;

  bool isSelected(int currentIndex) => index == currentIndex;
}

class AppNavigationBar extends StatelessWidget {
  const AppNavigationBar({
    super.key,
    required this.navigationShell,
  });

  final StatefulNavigationShell navigationShell;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      extendBody: true,
      body: Stack(
        alignment: Alignment.bottomCenter,
        children: [
          navigationShell,
          const _BottomBarBlur(),
        ],
      ),
      bottomNavigationBar: ClipRRect(
        borderRadius: const BorderRadius.only(
          topLeft: Radius.circular(50),
          topRight: Radius.circular(50),
        ),
        child: NavigationBar(
          height: 80,
          backgroundColor: Colors.black.withOpacity(0.7),
          selectedIndex: navigationShell.currentIndex,
          destinations: [
            _CustomBottomNavigationBarItem(
              icon: Icons.home,
              onPressed: () => _goBranch(BottomNavigationBarType.home.index),
              isSelected: BottomNavigationBarType.home
                  .isSelected(navigationShell.currentIndex),
            ),
            _CustomBottomNavigationBarItem(
              icon: Icons.search,
              onPressed: () => _goBranch(BottomNavigationBarType.search.index),
              isSelected: BottomNavigationBarType.search
                  .isSelected(navigationShell.currentIndex),
            ),
            _CenterCustomBottomNavigationBarItem(
              icon: Icons.add,
              // FIXME(onishi): 画面を作成していないため一旦空のコールバック指定。
              onPressed: () {},
            ),
            _CustomBottomNavigationBarItem(
              icon: Icons.notifications,
              // FIXME(onishi): 画面を作成していないため一旦空のコールバック指定。
              onPressed: () {},
              isSelected: BottomNavigationBarType.notification
                  .isSelected(navigationShell.currentIndex),
            ),
            _CustomBottomNavigationBarItem(
              icon: Icons.settings,
              // FIXME(onishi): 画面を作成していないため一旦空のコールバック指定。
              onPressed: () {},
              isSelected: BottomNavigationBarType.setting
                  .isSelected(navigationShell.currentIndex),
            ),
          ],
          onDestinationSelected: _goBranch,
        ),
      ),
    );
  }

  void _goBranch(int index) => navigationShell.goBranch(
        index,
        initialLocation: index == navigationShell.currentIndex,
      );
}

final class _CenterCustomBottomNavigationBarItem extends StatelessWidget {
  const _CenterCustomBottomNavigationBarItem({
    required this.onPressed,
    required this.icon,
  });

  final VoidCallback onPressed;
  final IconData icon;

  @override
  Widget build(BuildContext context) => Padding(
        padding: const EdgeInsets.only(top: 16),
        child: GestureDetector(
          onTap: onPressed,
          child: Container(
            padding: const EdgeInsets.all(1),
            decoration: BoxDecoration(
              border: Border.all(
                color: Colors.orange,
                width: 3,
              ),
              borderRadius: BorderRadius.circular(50),
            ),
            child: Container(
              margin: const EdgeInsets.all(2),
              padding: const EdgeInsets.all(15),
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(50),
                color: const Color.fromARGB(255, 36, 80, 99),
              ),
              child: Icon(
                icon,
                color: Colors.white,
              ),
            ),
          ),
        ),
      );
}

final class _CustomBottomNavigationBarItem extends StatelessWidget {
  const _CustomBottomNavigationBarItem({
    required this.onPressed,
    required this.icon,
    required this.isSelected,
  });

  final VoidCallback onPressed;
  final IconData icon;
  final bool isSelected;

  @override
  Widget build(BuildContext context) => Padding(
        padding: const EdgeInsets.only(top: 24),
        child: Column(
          children: [
            IconButton(
              onPressed: onPressed,
              icon: Icon(
                icon,
                size: 30,
                color: Colors.white,
              ),
            ),
            Visibility(
              visible: isSelected,
              child: Container(
                width: 5,
                height: 5,
                decoration: BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.circular(50),
                ),
              ).animate().fade(),
            ),
          ],
        ),
      );
}

final class _BottomBarBlur extends StatelessWidget {
  const _BottomBarBlur();

  @override
  Widget build(BuildContext context) => ClipRRect(
        borderRadius: const BorderRadius.only(
          topLeft: Radius.circular(50),
          topRight: Radius.circular(50),
        ),
        child: Stack(
          children: [
            const SizedBox(
              width: double.infinity,
              height: 114,
            ),
            Positioned.fill(
              child: ClipRect(
                child: BackdropFilter(
                  filter: ImageFilter.blur(sigmaX: 5, sigmaY: 5),
                  child: ColoredBox(
                    color: Colors.white.withOpacity(0.2),
                  ),
                ),
              ),
            ),
          ],
        ),
      );
}
