const CraftLuaEvents = {
    listeners: {},
    on(name, fn) {
        if (!this.listeners[name]) this.listeners[name] = [];
        this.listeners[name].push(fn);
    },
    emit(name, detail) {
        const list = this.listeners[name] || [];
        for (const fn of list) fn(detail);
    }
};