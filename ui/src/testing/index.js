import {act} from "react-dom/test-utils";

export async function updateComponent(component) {
  await act(async () => {
    await wait(0);
    component.update();
  });
}