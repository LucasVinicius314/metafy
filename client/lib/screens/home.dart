import 'package:flutter/material.dart';
import 'package:metafy/env.dart';
import 'package:web_socket_client/web_socket_client.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  static const route = '/';

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  final _wsChannel =
      WebSocket(Uri.parse('ws://${Env.apiAuthority}/api/connect'));

  var _sentAt = 0;

  @override
  void initState() {
    super.initState();

    _wsChannel.messages.listen((event) {
      final message = event as String;

      final now = DateTime.now().millisecondsSinceEpoch;

      print(message);

      print('${(now - _sentAt).toStringAsFixed(0)} ms');

      ScaffoldMessenger.of(context)
          .showSnackBar(SnackBar(content: Text(message)));
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          setState(() {
            _sentAt = DateTime.now().millisecondsSinceEpoch;
          });

          _wsChannel.send('hello world');
        },
      ),
    );
  }
}
