const CraftLuaSync = {
    intervalId: null,
    start(ms = 5000) {
        if (this.intervalId) clearInterval(this.intervalId);
        this.intervalId = setInterval(() => {
            CraftLuaEvents.emit("sync:tick", { time: Date.now() });
        }, ms);
    }
};