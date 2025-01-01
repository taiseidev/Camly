import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'environment_provider.g.dart';

enum Environment { dev, stg, prod }

@Riverpod(keepAlive: true)
Environment environment(Ref ref) => Environment.dev;
