import 'package:flutter/material.dart';
import 'package:camly_i18n/gen/strings.g.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

import 'core/providers/loading_notifier.dart';
import 'core/providers/navigator_key.dart';
import 'core/providers/scaffold_messenger_key.dart';
import 'core/navigation/app_router.dart';

final class App extends ConsumerWidget {
  const App({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final goRouter = ref.watch(goRouterProvider);

    return ScreenUtilInit(
      minTextAdapt: true,

      /// 画面サイズが確実に取得できるまで待つ。
      ensureScreenSize: true,
      builder: (context, _) => MaterialApp.router(
        routerConfig: goRouter,
        locale: TranslationProvider.of(context).flutterLocale,
        supportedLocales: AppLocaleUtils.supportedLocales,
        localizationsDelegates: GlobalMaterialLocalizations.delegates,
        debugShowCheckedModeBanner: false,
        scaffoldMessengerKey: ref.watch(scaffoldMessengerKeyProvider),
        builder: (context, child) => Consumer(
          builder: (context, ref, _) {
            final isLoading = ref.watch(loadingNotifierProvider);

            return Navigator(
              key: ref.watch(navigatorKeyProvider),
              // onDidRemovePage: (Page<Object?> page) => false,
              pages: [
                MaterialPage<Widget>(
                  child: Stack(
                    children: [
                      child!,
                      if (isLoading)
                        const ColoredBox(
                          color: Colors.black26,
                          child: Center(
                            child: CircularProgressIndicator(),
                          ),
                        ),
                    ],
                  ),
                ),
              ],
            );
          },
        ),
      ),
    );
  }
}
