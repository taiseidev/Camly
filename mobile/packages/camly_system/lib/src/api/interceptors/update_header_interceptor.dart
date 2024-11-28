import 'package:camly_system/src/storage/secure_storage.dart';
import 'package:dio/dio.dart';
import 'package:package_info_plus/package_info_plus.dart';

final class UpdateHeaderInterceptor extends Interceptor {
  const UpdateHeaderInterceptor();

  @override
  void onRequest(
    RequestOptions options,
    RequestInterceptorHandler handler,
  ) async {
    final accessToken = await SecureStorageKey.accessToken.read();
    final appVersion =
        await PackageInfo.fromPlatform().then((value) => value.version);

    options.headers['Content-Type'] = 'application/json';
    options.headers['AppVersion'] = appVersion;
    options.headers['Authorization'] = 'Bearer $accessToken';

    return handler.next(options);
  }
}
