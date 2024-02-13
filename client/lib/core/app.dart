import 'package:flutter/material.dart';
import 'package:metafy/screens/home.dart';

class App extends StatelessWidget {
  const App({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      routes: {
        HomeScreen.route: (context) => HomeScreen(),
      },
    );
  }
}
