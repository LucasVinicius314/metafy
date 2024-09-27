class AuthEvent {
  const AuthEvent();
}

class LoginAuthEvent extends AuthEvent {
  const LoginAuthEvent({
    required this.email,
    required this.password,
  });

  final String email;
  final String password;
}
