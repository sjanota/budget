import { act } from 'react-dom/test-utils';
import wait from 'waait';

export async function updateComponent(component) {
  await act(async () => {
    await wait(0);
    component.update();
  });
}
