/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { CurrencyAmountType } from './CurrencyAmountType';
import {
    CurrencyAmountTypeFromJSON,
    CurrencyAmountTypeFromJSONTyped,
    CurrencyAmountTypeToJSON,
} from './CurrencyAmountType';
import type { EventResourceNoteType } from './EventResourceNoteType';
import {
    EventResourceNoteTypeFromJSON,
    EventResourceNoteTypeFromJSONTyped,
    EventResourceNoteTypeToJSON,
} from './EventResourceNoteType';
import type { ResourceId } from './ResourceId';
import {
    ResourceIdFromJSON,
    ResourceIdFromJSONTyped,
    ResourceIdToJSON,
} from './ResourceId';
import type { ResourceType } from './ResourceType';
import {
    ResourceTypeFromJSON,
    ResourceTypeFromJSONTyped,
    ResourceTypeToJSON,
} from './ResourceType';

/**
 * Event resource information.
 * @export
 * @interface CateringEventResourceType
 */
export interface CateringEventResourceType {
    /**
     * 
     * @type {ResourceId}
     * @memberof CateringEventResourceType
     */
    resourceId?: ResourceId;
    /**
     * 
     * @type {ResourceType}
     * @memberof CateringEventResourceType
     */
    resourceType?: ResourceType;
    /**
     * Resource name, this could be a resource item or a menu.
     * @type {string}
     * @memberof CateringEventResourceType
     */
    resourceName?: string;
    /**
     * Setup style for the function space.
     * @type {string}
     * @memberof CateringEventResourceType
     */
    setupCode?: string;
    /**
     * Item attribute name for non Food and Beverage item.
     * @type {string}
     * @memberof CateringEventResourceType
     */
    itemAttribute?: string;
    /**
     * The required quantity of the resource.
     * @type {number}
     * @memberof CateringEventResourceType
     */
    quantity?: number;
    /**
     * The number of charged hours for the resource.
     * @type {number}
     * @memberof CateringEventResourceType
     */
    hours?: number;
    /**
     * The number of charged persons for the resource.
     * @type {number}
     * @memberof CateringEventResourceType
     */
    persons?: number;
    /**
     * The attendees/quantity of an item INCLUDED in a package price.
     * @type {number}
     * @memberof CateringEventResourceType
     */
    quantityIncluded?: number;
    /**
     * The attendees/quantity charged extra, EXCLUDED from the package price.
     * @type {number}
     * @memberof CateringEventResourceType
     */
    quantityExcluded?: number;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof CateringEventResourceType
     */
    unitPrice?: CurrencyAmountType;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof CateringEventResourceType
     */
    revenue?: CurrencyAmountType;
    /**
     * Rate Code that is attached to the function spaces and non Food and Beverage items
     * @type {string}
     * @memberof CateringEventResourceType
     */
    rentalCode?: string;
    /**
     * The menu discount percentage.
     * @type {number}
     * @memberof CateringEventResourceType
     */
    discount?: number;
    /**
     * Contains event resource comment information.
     * @type {Array<EventResourceNoteType>}
     * @memberof CateringEventResourceType
     */
    eventResourceNotes?: Array<EventResourceNoteType>;
    /**
     * Sell Start Date of resource of type Menu.
     * @type {Date}
     * @memberof CateringEventResourceType
     */
    sellStartDate?: Date;
    /**
     * Sell End Date of resource of type Menu.
     * @type {Date}
     * @memberof CateringEventResourceType
     */
    sellEndDate?: Date;
    /**
     * Event Start Date of resource of type Menu.
     * @type {Date}
     * @memberof CateringEventResourceType
     */
    eventStartDate?: Date;
    /**
     * Event End Date of resource of type Menu.
     * @type {Date}
     * @memberof CateringEventResourceType
     */
    eventEndDate?: Date;
    /**
     * Indicates that this resource is consumption based in a catering package
     * @type {boolean}
     * @memberof CateringEventResourceType
     */
    consumptionBased?: boolean;
    /**
     * Indicates that this resource is a MultiChoice Menu or not
     * @type {boolean}
     * @memberof CateringEventResourceType
     */
    multiChoice?: boolean;
    /**
     * Hotel code which the event resource belong to.
     * @type {string}
     * @memberof CateringEventResourceType
     */
    hotelId?: string;
    /**
     * The order of the resource in the Banquet Event Order within their specific resource group.
     * @type {string}
     * @memberof CateringEventResourceType
     */
    order?: string;
    /**
     * Flag that indicates if at least one menu item has a discount.
     * @type {boolean}
     * @memberof CateringEventResourceType
     */
    hasDiscountMenuItem?: boolean;
    /**
     * Flag that indicates the resource has notes associated with it.
     * @type {boolean}
     * @memberof CateringEventResourceType
     */
    hasNotes?: boolean;
    /**
     * Indicates that item must be ordered externally.
     * @type {boolean}
     * @memberof CateringEventResourceType
     */
    external?: boolean;
    /**
     * Indicates that this resource item is a custom item added for this event.
     * @type {boolean}
     * @memberof CateringEventResourceType
     */
    custom?: boolean;
    /**
     * Indicates that this resource is included in a catering package.
     * @type {boolean}
     * @memberof CateringEventResourceType
     */
    packaged?: boolean;
}

/**
 * Check if a given object implements the CateringEventResourceType interface.
 */
export function instanceOfCateringEventResourceType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CateringEventResourceTypeFromJSON(json: any): CateringEventResourceType {
    return CateringEventResourceTypeFromJSONTyped(json, false);
}

export function CateringEventResourceTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CateringEventResourceType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'resourceId': !exists(json, 'resourceId') ? undefined : ResourceIdFromJSON(json['resourceId']),
        'resourceType': !exists(json, 'resourceType') ? undefined : ResourceTypeFromJSON(json['resourceType']),
        'resourceName': !exists(json, 'resourceName') ? undefined : json['resourceName'],
        'setupCode': !exists(json, 'setupCode') ? undefined : json['setupCode'],
        'itemAttribute': !exists(json, 'itemAttribute') ? undefined : json['itemAttribute'],
        'quantity': !exists(json, 'quantity') ? undefined : json['quantity'],
        'hours': !exists(json, 'hours') ? undefined : json['hours'],
        'persons': !exists(json, 'persons') ? undefined : json['persons'],
        'quantityIncluded': !exists(json, 'quantityIncluded') ? undefined : json['quantityIncluded'],
        'quantityExcluded': !exists(json, 'quantityExcluded') ? undefined : json['quantityExcluded'],
        'unitPrice': !exists(json, 'unitPrice') ? undefined : CurrencyAmountTypeFromJSON(json['unitPrice']),
        'revenue': !exists(json, 'revenue') ? undefined : CurrencyAmountTypeFromJSON(json['revenue']),
        'rentalCode': !exists(json, 'rentalCode') ? undefined : json['rentalCode'],
        'discount': !exists(json, 'discount') ? undefined : json['discount'],
        'eventResourceNotes': !exists(json, 'eventResourceNotes') ? undefined : ((json['eventResourceNotes'] as Array<any>).map(EventResourceNoteTypeFromJSON)),
        'sellStartDate': !exists(json, 'sellStartDate') ? undefined : (new Date(json['sellStartDate'])),
        'sellEndDate': !exists(json, 'sellEndDate') ? undefined : (new Date(json['sellEndDate'])),
        'eventStartDate': !exists(json, 'eventStartDate') ? undefined : (new Date(json['eventStartDate'])),
        'eventEndDate': !exists(json, 'eventEndDate') ? undefined : (new Date(json['eventEndDate'])),
        'consumptionBased': !exists(json, 'consumptionBased') ? undefined : json['consumptionBased'],
        'multiChoice': !exists(json, 'multiChoice') ? undefined : json['multiChoice'],
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'order': !exists(json, 'order') ? undefined : json['order'],
        'hasDiscountMenuItem': !exists(json, 'hasDiscountMenuItem') ? undefined : json['hasDiscountMenuItem'],
        'hasNotes': !exists(json, 'hasNotes') ? undefined : json['hasNotes'],
        'external': !exists(json, 'external') ? undefined : json['external'],
        'custom': !exists(json, 'custom') ? undefined : json['custom'],
        'packaged': !exists(json, 'packaged') ? undefined : json['packaged'],
    };
}

export function CateringEventResourceTypeToJSON(value?: CateringEventResourceType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'resourceId': ResourceIdToJSON(value.resourceId),
        'resourceType': ResourceTypeToJSON(value.resourceType),
        'resourceName': value.resourceName,
        'setupCode': value.setupCode,
        'itemAttribute': value.itemAttribute,
        'quantity': value.quantity,
        'hours': value.hours,
        'persons': value.persons,
        'quantityIncluded': value.quantityIncluded,
        'quantityExcluded': value.quantityExcluded,
        'unitPrice': CurrencyAmountTypeToJSON(value.unitPrice),
        'revenue': CurrencyAmountTypeToJSON(value.revenue),
        'rentalCode': value.rentalCode,
        'discount': value.discount,
        'eventResourceNotes': value.eventResourceNotes === undefined ? undefined : ((value.eventResourceNotes as Array<any>).map(EventResourceNoteTypeToJSON)),
        'sellStartDate': value.sellStartDate === undefined ? undefined : (value.sellStartDate.toISOString().substr(0,10)),
        'sellEndDate': value.sellEndDate === undefined ? undefined : (value.sellEndDate.toISOString().substr(0,10)),
        'eventStartDate': value.eventStartDate === undefined ? undefined : (value.eventStartDate.toISOString().substr(0,10)),
        'eventEndDate': value.eventEndDate === undefined ? undefined : (value.eventEndDate.toISOString().substr(0,10)),
        'consumptionBased': value.consumptionBased,
        'multiChoice': value.multiChoice,
        'hotelId': value.hotelId,
        'order': value.order,
        'hasDiscountMenuItem': value.hasDiscountMenuItem,
        'hasNotes': value.hasNotes,
        'external': value.external,
        'custom': value.custom,
        'packaged': value.packaged,
    };
}

