const SimpleStorage = artifacts.require("SimpleStorage");

contract("SimpleStorage", accounts => {
    let simpleStorage;

    beforeEach(async () => {
        simpleStorage = await SimpleStorage.new();
    });

    it("should store and retrieve a value", async () => {
        const value = 42;
        await simpleStorage.set(value);
        const result = await simpleStorage.get();
        assert.equal(result.toNumber(), value, "The value was not stored correctly");
    });
});