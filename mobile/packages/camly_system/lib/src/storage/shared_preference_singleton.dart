import 'package:shared_preferences/shared_preferences.dart';

final class SharedPreferencesSingleton {
  static final SharedPreferencesSingleton _instance =
      SharedPreferencesSingleton._internal();

  late final SharedPreferences _preferences;

  SharedPreferencesSingleton._internal();

  factory SharedPreferencesSingleton() => _instance;

  Future<void> initialize() async =>
      _preferences = await SharedPreferences.getInstance();

  SharedPreferences get instance => _preferences;
}
