import React, { useEffect, useState } from 'react';
import ReactDOM from 'react-dom';
import './sb-admin-2.css';
import './index.css';
import App from './components/App/App';
import * as serviceWorker from './serviceWorker';
import { BrowserRouter } from 'react-router-dom';
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';
import { ApolloProvider } from '@apollo/react-hooks';
import createClient from './apollo';
import { Auth0Provider, Auth0Context, useAuth0 } from './react-auth0-spa';
import config from './auth_config.json';

// A function that routes the user to the right place
// after login
const onRedirectCallback = appState => {
  window.history.replaceState(
    {},
    document.title,
    appState && appState.targetUrl
      ? appState.targetUrl
      : window.location.pathname
  );
};

const ProdAuthorizationProvider = ({ children }) => (
  <Auth0Provider
    domain={config.domain}
    client_id={config.clientId}
    redirect_uri={window.location.origin}
    onRedirectCallback={onRedirectCallback}
    audience={config.audience}
  >
    {children}
  </Auth0Provider>
);

const DevAuthorizationProvider = ({ children }) => (
  <Auth0Context.Provider
    value={{
      isAuthenticated: true,
      loading: false,
      loginWithRedirect: () => {},
      user: {
        name: 'Valerie Luna',
        picture: 'https://source.unsplash.com/QAB-WJcbgJk/60x60',
      },
    }}
  >
    {children}
  </Auth0Context.Provider>
);

const authDisabled = process.env.REACT_APP_INSECURE_AUTH_DISABLED;
const AuthorizationProvider =
  authDisabled !== 'true'
    ? ProdAuthorizationProvider
    : DevAuthorizationProvider;

function Apollo({ children }) {
  const {
    isAuthenticated,
    loading,
    loginWithRedirect,
    getTokenSilently,
  } = useAuth0();
  const [token, setToken] = useState();

  useEffect(() => {
    if (loading) {
      return;
    }
    if (!isAuthenticated) {
      loginWithRedirect({});
      return;
    }

    getTokenSilently().then(setToken);
  }, [isAuthenticated, loginWithRedirect, loading, getTokenSilently]);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (!isAuthenticated || !token) {
    return <div />;
  }

  return (
    <ApolloProvider client={createClient(token)}>{children}</ApolloProvider>
  );
}

ReactDOM.render(
  <AuthorizationProvider>
    <Apollo>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </Apollo>
  </AuthorizationProvider>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
