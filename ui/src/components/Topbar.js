import React, { useEffect } from 'react';
import TemplateTopbar from './template/Topbar/Topbar';
import { TopbarMenu } from './template/Topbar/TopbarMenu';
import TopbarUser from './template/Topbar/TopbarUser';
import TopbarBudgetSwitcher from './TopbarBudgetSwitcher';
import { useAuth0 } from '../react-auth0-spa';
import { useDictionary } from './template/Utilities/Lang';

export default function Topbar() {
  const { user, logout } = useAuth0();
  const { topbar } = useDictionary();
  return (
    <TemplateTopbar
      faIconContextMinified="search"
      renderContext={() => <TopbarBudgetSwitcher />}
      renderUser={() => (
        <TopbarUser
          name={user.name}
          pictureUrl={user.picture}
          logout={logout}
        />
      )}
      renderMenus={() => (
        <>
          <TopbarMenu
            name={topbar.alertsLabel}
            faIcon="bell"
            counter={0}
          ></TopbarMenu>
          <TopbarMenu
            name={topbar.messagesLabel}
            faIcon="envelope"
            counter={3}
          ></TopbarMenu>
        </>
      )}
    />
  );
}
