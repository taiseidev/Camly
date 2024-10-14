import 'package:flutter/material.dart';

final class TopPage extends StatelessWidget {
  const TopPage({super.key});

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text('アプリのエントリーポイント'),
          ],
        ),
      ),
    );
  }
}
