import { useState } from 'react';
import SignIn from './SignIn';
import SignUp from './SignUp';
import SignOut from './SignOut';

type AuthStatus = 'signedOut' | 'signedIn';
type AuthView = 'signin' | 'signup';

const AuthPage = () => {
  const [authStatus, setAuthStatus] = useState<AuthStatus>('signedOut');
  const [currentView, setCurrentView] = useState<AuthView>('signin');

  const handleSignInSuccess = () => {
    setAuthStatus('signedIn');
  };

  const handleSignUpSuccess = () => {
    setCurrentView('signin');
  };

  const handleSignOutSuccess = () => {
    setAuthStatus('signedOut');
    setCurrentView('signin');
  };

  const toggleView = () => {
    setCurrentView(currentView === 'signin' ? 'signup' : 'signin');
  };

  if (authStatus === 'signedIn') {
    return (
      <div>
        <h1>Welcome to the App</h1>
        <p>You are signed in!</p>
        <SignOut onSuccess={handleSignOutSuccess} />
      </div>
    );
  }

  return (
    <div>
      <h1>Authentication</h1>
      {currentView === 'signin' ? (
        <>
          <SignIn onSuccess={handleSignInSuccess} />
          <p>
            Don't have an account?{' '}
            <button onClick={toggleView}>Sign Up</button>
          </p>
        </>
      ) : (
        <>
          <SignUp onSuccess={handleSignUpSuccess} />
          <p>
            Already have an account?{' '}
            <button onClick={toggleView}>Sign In</button>
          </p>
        </>
      )}
    </div>
  );
};

export default AuthPage; 