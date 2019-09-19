import { configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import 'jest-enzyme';
import wait from 'waait';

configure({ adapter: new Adapter() });
window.wait = wait;