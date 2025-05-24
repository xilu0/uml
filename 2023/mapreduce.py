import matplotlib.pyplot as plt
import matplotlib.patches as mpatches
from matplotlib.offsetbox import AnnotationBbox, TextArea

fig, ax = plt.subplots()

# Hide axes
ax.axis('off')

# Function to create a UML class box
def draw_class(ax, class_name, attributes, methods, xy):
    # Draw the main box
    height = 0.2 * (2 + len(attributes) + len(methods))
    width = 0.2
    main_box = mpatches.FancyBboxPatch(
        xy, width, height, boxstyle="round,pad=0.05", edgecolor="black", facecolor="white"
    )
    ax.add_patch(main_box)

    # Add the class name
    text_offset = 0.1
    class_text = TextArea(class_name, textprops=dict(color="black", size=10, ha='center'))
    ab = AnnotationBbox(class_text, (xy[0] + width / 2, xy[1] + height - text_offset), frameon=False)
    ax.add_artist(ab)

    # Add attributes and methods
    current_height = xy[1] + height - 2 * text_offset
    for attr in attributes:
        attr_text = TextArea(attr, textprops=dict(color="black", size=8))
        ab = AnnotationBbox(attr_text, (xy[0] + width / 2, current_height), frameon=False)
        ax.add_artist(ab)
        current_height -= text_offset

    for method in methods:
        method_text = TextArea(method, textprops=dict(color="black", size=8))
        ab = AnnotationBbox(method_text, (xy[0] + width / 2, current_height), frameon=False)
        ax.add_artist(ab)
        current_height -= text_offset

# Draw UML classes
draw_class(ax, "Main", [], ["+main()"], (0, 0))
draw_class(ax, "Map", [], ["+Map(word: string, ch: channel)"], (0.5, 0))
draw_class(ax, "Reduce", [], ["+Reduce(frequencies: []map[string]int): map[string]int"], (1, 0))

# Set figure size and title
fig.set_size_inches(7, 3)
plt.title("MapReduce Implementation in Go - UML Diagram")

plt.show()
