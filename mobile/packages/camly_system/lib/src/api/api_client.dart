import 'package:dio/dio.dart';

import 'api_exception.dart';
import 'interceptors/update_header_interceptor.dart';
import 'interceptors/update_tokens_interceptor.dart';

final class ApiClient {
  ApiClient(String baseUrl)
      : _dio = Dio(
          BaseOptions(
            baseUrl: baseUrl,
            connectTimeout: const Duration(milliseconds: 5000),
            receiveTimeout: const Duration(milliseconds: 5000),
          ),
        ) {
    _dio.interceptors.addAll(
      [
        LogInterceptor(),
        const UpdateHeaderInterceptor(),
        const UpdateTokensInterceptor(),
      ],
    );
  }

  final Dio _dio;

  /// GETリクエスト
  Future<T> get<T>(
    String endpoint, {
    Map<String, dynamic>? params,
  }) async {
    try {
      final response = await _dio.get(endpoint, queryParameters: params);
      return response.data;
    } on DioException catch (e) {
      throw ApiException.fromDioError(e);
    } catch (e) {
      throw ApiException(message: e.toString());
    }
  }

  /// POSTリクエスト
  Future<T> post<T>(
    String endpoint, {
    Map<String, dynamic>? data,
  }) async {
    try {
      final response = await _dio.post(endpoint, data: data);
      return response.data;
    } on DioException catch (e) {
      throw ApiException.fromDioError(e);
    } catch (e) {
      throw ApiException(message: e.toString());
    }
  }
}
