import 'package:flutter/material.dart';
import 'package:metafy/blocs/main/main_bloc.dart';
import 'package:metafy/blocs/main/main_event.dart';
import 'package:metafy/utils/constants.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  static const route = '/';

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  final _mainBloc = MainBloc();

  @override
  void initState() {
    super.initState();

    _mainBloc.add(ConnectMainEvent(token: 'fh98h9g854high'));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text(Constants.appName)),
    );
  }
}
