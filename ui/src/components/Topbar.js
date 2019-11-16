import React from 'react';
import TemplateTopbar from './template/Topbar/Topbar';
import TopbarAlert from './template/Topbar/TopbarAlert';
import TopbarMessage from './template/Topbar/TopbarMessage';
import { TopbarMenu } from './template/Topbar/TopbarMenu';
import TopbarUser from './template/Topbar/TopbarUser';
import TopbarBudgetSwitcher from './TopbarBudgetSwitcher';
import { useAuth0 } from '../react-auth0-spa';

export default function Topbar() {
  const { user, logout } = useAuth0();
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
          <TopbarMenu name="Alerts center" faIcon="bell" counter={0}>
            <TopbarAlert
              highlighted={true}
              faIcon="file-alt"
              date="December 12, 2019"
              text="A new monthly report is ready to download!"
            />
            <TopbarAlert
              faIcon="donate"
              date="December 7, 2019"
              variant="success"
              text="$290.29 has been deposited into your account!"
            />
            <TopbarAlert
              faIcon="exclamation-triangle"
              date="December 2, 2019"
              variant="warning"
              text="Spending Alert: We've noticed unusually high spending for your
                account."
            />
          </TopbarMenu>
          <TopbarMenu name="Message Center" faIcon="envelope" counter={3}>
            <TopbarMessage
              highlighted={true}
              imgSrc="https://source.unsplash.com/fn_BT9fwg_E/60x60"
              variant="success"
              text="Hi there! I am wondering if you can help me with a problem I've been
                  having."
              author="Emily Fowler"
              when="58m"
            />
            <TopbarMessage
              highlighted={true}
              imgSrc="https://source.unsplash.com/AU4VPcFN4LE/60x60"
              text="I have the photos that you ordered last month, how would you
                  like them sent to you?"
              author="Jae Chun"
              when="1d"
            />
            <TopbarMessage
              highlighted={true}
              imgSrc="https://source.unsplash.com/CS2uCrpNzJY/60x60"
              text="Last month's report looks great, I am very happy with the
                  progress so far, keep up the good work!"
              author="Morgan Alvarez"
              when="2d"
            />
            <TopbarMessage
              imgSrc="https://source.unsplash.com/Mv9hjnEUHR4/60x60"
              text="Am I a good boy? The reason I ask is because someone told me
                  that people say this to all dogs, even if they aren't good..."
              author="Chicken the Dog"
              when="2w"
            />
          </TopbarMenu>
        </>
      )}
    />
  );
}
