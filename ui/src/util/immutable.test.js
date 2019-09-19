import {addToList, removeFromList, removeFromListByID, replaceOnList, replaceOnListByID} from "./immutable";

describe('addToList', () => {
  describe('when element is added', () => {
    const oldList = [1, 2, 3];
    const el = 4;

    const newList = addToList(oldList, el);

    it('it is appended', () => {
      expect(newList).toEqual([1, 2, 3, 4]);
    });

    it('old list is preserved', () => {
      expect(oldList).toEqual([1, 2, 3]);
      expect(oldList).not.toBe(newList);
    });
  });
});

describe('removeFromList', () => {
    describe('when element is removed', () => {
      const oldList = [1, 2, 3];
      const el = 2;

      const newList = removeFromList(oldList, el);

      it('it is removed', () => {
        expect(newList).toEqual([1, 3]);
      });

      it('old list is preserved', () => {
        expect(oldList).toEqual([1, 2, 3]);
        expect(oldList).not.toBe(newList);
      });
    });

    describe('when element is not on the list', () => {
      const oldList = [1, 2, 3];
      const el = 5;

      const newList = removeFromList(oldList, el);

      it('ignores the operation', () => {
        expect(newList).toBe(oldList)
      });
    });
});

describe('removeFromListByID', () => {
  describe('when element is removed', () => {
    const oldList = [{id: 1}, {id: 2}, {id: 3}];
    const el = 2;

    const newList = removeFromListByID(oldList, el);

    it('it is removed', () => {
      expect(newList).toEqual([{id: 1}, {id: 3}]);
    });

    it('old list is preserved', () => {
      expect(oldList).toEqual([{id: 1}, {id: 2}, {id: 3}]);
      expect(oldList).not.toBe(newList);
    });
  });

  describe('when element is not on the list', () => {
    const oldList = [{id: 1}, {id: 2}, {id: 3}];
    const el = 5;

    const newList = removeFromListByID(oldList, el);

    it('ignores the operation', () => {
      expect(newList).toBe(oldList)
    });
  });
});

describe('replaceOnList', () => {
  describe('when element is replaced', () => {
    const oldList = [1, 2, 3];
    const idx = 1;
    const el = 5;

    const newList = replaceOnList(oldList, idx, el);

    it('it is removed', () => {
      expect(newList).toEqual([1, 5, 3]);
    });

    it('old list is preserved', () => {
      expect(oldList).toEqual([1, 2, 3]);
      expect(oldList).not.toBe(newList);
    });
  });

  describe('when index is outside the list', () => {
    const oldList = [1, 2, 3];
    const idx = 5;
    const el = 5;

    const newList = replaceOnList(oldList, idx, el);

    it('ignores the operation', () => {
      expect(newList).toBe(oldList)
    });
  });

  describe('when index is not on the list', () => {
    const oldList = [1, 2, 3];
    const idx = -1;
    const el = 5;

    const newList = replaceOnList(oldList, idx, el);

    it('ignores the operation', () => {
      expect(newList).toBe(oldList)
    });
  });
});

describe('replaceOnListByID', () => {
  describe('when element is present', () => {
    const oldList = [{id: 1}, {id: 2, a: "b"}, {id: 3}];
    const el = {id: 2, a: "c"};

    const newList = replaceOnListByID(oldList, el);

    it('it is replaced', () => {
      expect(newList).toEqual([{id: 1}, {id: 2, a: "c"}, {id: 3}]);
    });

    it('old list is preserved', () => {
      expect(oldList).toEqual([{id: 1}, {id: 2, a: "b"}, {id: 3}]);
      expect(oldList).not.toBe(newList);
    });
  });

  describe('when element is not on the list', () => {
    const oldList = [{id: 1}, {id: 2}, {id: 3}];
    const el = {id: 5};

    const newList = removeFromListByID(oldList, el);

    it('ignores the operation', () => {
      expect(newList).toBe(oldList)
    });
  });
});