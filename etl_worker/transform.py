def count_events(rows):
    # rows: [('purchase',), ('click',)] gibi
    counts = {}
    for r in rows:
        event_name = r[0] # Tuple içindeki ilk eleman
        if event_name in counts:
            counts[event_name] += 1
        else:
            counts[event_name] = 1
    return counts # Örn: {'purchase': 2, 'click': 1} gibi