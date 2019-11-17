import { useAuth0 } from '../react-auth0-spa';
import React from 'react';

export function Beta({ children }) {
  const { token } = useAuth0();
  const claimsString = token.split('.')[1];
  const claims = JSON.parse(atob(claimsString));
  const scope = claims.scope.split(' ');
  const isBetaUser = scope.some(v => v === 'beta');
  if (isBetaUser) {
    return children;
  }
  return <div>This conent is available only to beta users</div>;
}
