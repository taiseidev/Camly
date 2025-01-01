import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'use_mock_provider.g.dart';

@Riverpod(keepAlive: true)
bool useMock(Ref ref) => false;
