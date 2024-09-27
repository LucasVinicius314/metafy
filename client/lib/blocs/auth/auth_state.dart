class AuthState {
  const AuthState();
}

class LoadingAuthState extends AuthState {
  const LoadingAuthState();
}

// Login.

class LoginDoneAuthState extends AuthState {
  const LoginDoneAuthState({
    required this.token,
  });

  final String token;
}

class LoginErrorAuthState extends AuthState {
  const LoginErrorAuthState();
}
