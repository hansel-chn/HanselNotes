UPDATE table
SET conf_value =
        CASE
            WHEN conf_name = ? THEN
                ?
            WHEN conf_name = ? THEN
                ?
            WHEN conf_name = ? THEN
                ?
            ELSE conf_value END;