import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:metafy/blocs/main/main_event.dart';
import 'package:metafy/blocs/main/main_state.dart';
import 'package:metafy/env.dart';
import 'package:web_socket_client/web_socket_client.dart';

class MainBloc extends Bloc<MainEvent, MainState> {
  MainBloc() : super(const DisconnectedMainState()) {
    on<ConnectMainEvent>((event, emit) {
      final wsChannel = WebSocket(
        Uri.parse('ws://${Env.apiAuthority}/api/connect'),
        timeout: const Duration(days: 1),
        protocols: ['token-${event.token}'],
      );

      wsChannel.connection.listen((connectionState) {
        if (connectionState is Connecting) {
          emit(const ConnectingMainState());
        } else if (connectionState is Connected) {
          emit(const ConnectedMainState());
        } else if (connectionState is Reconnecting) {
          emit(const ConnectingMainState());
        } else if (connectionState is Reconnected) {
          emit(const ConnectedMainState());
        } else if (connectionState is Disconnecting) {
          emit(const ConnectingMainState());
        } else if (connectionState is Disconnected) {
          emit(const DisconnectedMainState());
        }
      });

      wsChannel.messages.listen((event) {
        final message = event as String;

        print(message);
      });

      _wsChannel = wsChannel;
    });
  }

  WebSocket? _wsChannel;
}
