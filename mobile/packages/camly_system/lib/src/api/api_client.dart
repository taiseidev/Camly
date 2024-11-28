import 'package:dio/dio.dart';

import 'api_exception.dart';
import 'interceptors/update_header_interceptor.dart';
import 'result.dart';

final class ApiClient {
  ApiClient(String baseUrl)
      : _dio = Dio(BaseOptions(
          baseUrl: baseUrl,
          connectTimeout: Duration(microseconds: 5000),
          receiveTimeout: Duration(microseconds: 3000),
        )) {
    this._dio.interceptors.addAll(
      [
        LogInterceptor(),
        UpdateHeaderInterceptor(),
      ],
    );
  }

  final Dio _dio;

  /// GETリクエスト
  Future<Result<T>> get<T>(
    String endpoint, {
    Map<String, dynamic>? params,
  }) async {
    try {
      final response = await _dio.get(endpoint, queryParameters: params);
      return Success(response.data);
    } on DioException catch (e) {
      return Failure(ApiException.fromDioError(e));
    } catch (e) {
      return Failure(ApiException(message: e.toString()));
    }
  }

  /// POSTリクエスト
  Future<Result<T>> post<T>(
    String endpoint, {
    Map<String, dynamic>? data,
  }) async {
    try {
      final response = await _dio.post(endpoint, data: data);
      return Success(response.data);
    } on DioException catch (e) {
      return Failure(ApiException.fromDioError(e));
    } catch (e) {
      return Failure(ApiException(message: e.toString()));
    }
  }
}
