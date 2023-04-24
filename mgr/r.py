import os
l = os.listdir()
# l = ["waybill.go"]
for filename in l:
    data_new = None
    with open(filename, "r") as fr:
        data = fr.read()
        start = data.find("type")
        end = data.find("struct")
        sss = data[start:end].strip()
        s = sss.split(" ")
        if len(s) > 2:
            continue
        old_mgr = s[-1]
        if old_mgr.startswith("_Es"):
            new_mgr = old_mgr.strip("_Es")
            data_new = data.replace(old_mgr, new_mgr)
            data_new = data_new.replace("Es"+new_mgr, "New"+new_mgr)
            data_new = data_new.replace("*gorm.DB", "db.Repo", 1)
            data_new = data_new.replace("DB", "rdb")
            data_new = data_new.replace("db.Table", "db.GetDbR().Table")
            data_new = data_new.replace("gorm.rdb", "gorm.DB")
    if data_new is not None:
        prefix, postfix = filename.split(".")
        new_file_name = prefix+"_new." + postfix
        with open(filename, "w") as fw:
            fw.write(data_new)