import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:metafy/blocs/auth/auth_event.dart';
import 'package:metafy/blocs/auth/auth_state.dart';
import 'package:metafy/env.dart';
import 'package:web_socket_client/web_socket_client.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  AuthBloc() : super(const LoadingAuthState()) {
    on<LoginAuthEvent>((event, emit) {
      // TODO: fix, api
    });
  }

  WebSocket? _wsChannel;
}
