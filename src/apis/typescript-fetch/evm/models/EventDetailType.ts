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
import type { BookingStatusType } from './BookingStatusType';
import {
    BookingStatusTypeFromJSON,
    BookingStatusTypeFromJSONTyped,
    BookingStatusTypeToJSON,
} from './BookingStatusType';
import type { CateringEventLinkType } from './CateringEventLinkType';
import {
    CateringEventLinkTypeFromJSON,
    CateringEventLinkTypeFromJSONTyped,
    CateringEventLinkTypeToJSON,
} from './CateringEventLinkType';
import type { CateringEventsAttendeesType } from './CateringEventsAttendeesType';
import {
    CateringEventsAttendeesTypeFromJSON,
    CateringEventsAttendeesTypeFromJSONTyped,
    CateringEventsAttendeesTypeToJSON,
} from './CateringEventsAttendeesType';
import type { CateringStatusTypeType } from './CateringStatusTypeType';
import {
    CateringStatusTypeTypeFromJSON,
    CateringStatusTypeTypeFromJSONTyped,
    CateringStatusTypeTypeToJSON,
} from './CateringStatusTypeType';
import type { DateTimeSpanType } from './DateTimeSpanType';
import {
    DateTimeSpanTypeFromJSON,
    DateTimeSpanTypeFromJSONTyped,
    DateTimeSpanTypeToJSON,
} from './DateTimeSpanType';
import type { TranslationTextType60 } from './TranslationTextType60';
import {
    TranslationTextType60FromJSON,
    TranslationTextType60FromJSONTyped,
    TranslationTextType60ToJSON,
} from './TranslationTextType60';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Pertain Information about the Event Details
 * @export
 * @interface EventDetailType
 */
export interface EventDetailType {
    /**
     * 
     * @type {TranslationTextType60}
     * @memberof EventDetailType
     */
    eventName?: TranslationTextType60;
    /**
     * Type of event.
     * @type {string}
     * @memberof EventDetailType
     */
    eventType?: string;
    /**
     * 
     * @type {DateTimeSpanType}
     * @memberof EventDetailType
     */
    eventTimeSpan?: DateTimeSpanType;
    /**
     * 
     * @type {BookingStatusType}
     * @memberof EventDetailType
     */
    eventStatus?: BookingStatusType;
    /**
     * 
     * @type {BookingStatusType}
     * @memberof EventDetailType
     */
    waitlistReturnStatus?: BookingStatusType;
    /**
     * 
     * @type {CateringStatusTypeType}
     * @memberof EventDetailType
     */
    cateringStatusType?: CateringStatusTypeType;
    /**
     * 
     * @type {CateringEventsAttendeesType}
     * @memberof EventDetailType
     */
    attendees?: CateringEventsAttendeesType;
    /**
     * 
     * @type {UniqueIDType}
     * @memberof EventDetailType
     */
    masterEventId?: UniqueIDType;
    /**
     * It is the group details to display on the outside of a meeting functionSpaceDetails.
     * @type {string}
     * @memberof EventDetailType
     */
    doorcard?: string;
    /**
     * 
     * @type {CateringEventLinkType}
     * @memberof EventDetailType
     */
    eventLink?: CateringEventLinkType;
    /**
     * Inactivation date of the event.
     * @type {Date}
     * @memberof EventDetailType
     */
    inactiveDate?: Date;
    /**
     * Package Id of the Event.
     * @type {number}
     * @memberof EventDetailType
     */
    packageId?: number;
    /**
     * Package Name of the Event.
     * @type {string}
     * @memberof EventDetailType
     */
    packageName?: string;
    /**
     * Indicates whether event is master event.
     * @type {boolean}
     * @memberof EventDetailType
     */
    masterEvent?: boolean;
    /**
     * Flag to indicate if the event is a sub event.
     * @type {boolean}
     * @memberof EventDetailType
     */
    subEvent?: boolean;
    /**
     * Flag to indicate if the event is part of a package.
     * @type {boolean}
     * @memberof EventDetailType
     */
    packageEvent?: boolean;
    /**
     * Flag to indicate if the event is booked as a backup for another event.
     * @type {boolean}
     * @memberof EventDetailType
     */
    alternateEvent?: boolean;
    /**
     * Indicates whether event is wait listed.
     * @type {boolean}
     * @memberof EventDetailType
     */
    waitlisted?: boolean;
    /**
     * Indicates whether event dates and functionSpaceDetails are move able.
     * @type {boolean}
     * @memberof EventDetailType
     */
    notMoveable?: boolean;
    /**
     * Indicates whether the event is expected to be noisy and might possibly disturb other events.
     * @type {boolean}
     * @memberof EventDetailType
     */
    loudEvent?: boolean;
    /**
     * Indicates whether the events has postings.
     * @type {boolean}
     * @memberof EventDetailType
     */
    hasPostings?: boolean;
    /**
     * Indicates whether display the doorcard information on the Event Overview report.
     * @type {boolean}
     * @memberof EventDetailType
     */
    displayDoorcard?: boolean;
    /**
     * Indicates if catering event's spaces are deducted from inventory
     * @type {boolean}
     * @memberof EventDetailType
     */
    cateringDeductInventory?: boolean;
    /**
     * Flag to indicate if the event is shareable with other events.
     * @type {boolean}
     * @memberof EventDetailType
     */
    eventShared?: boolean;
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof EventDetailType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof EventDetailType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof EventDetailType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof EventDetailType
     */
    lastModifierId?: string;
    /**
     * Date an item will be purged from a database (e.g., from a live database to an archive).
     * @type {Date}
     * @memberof EventDetailType
     */
    purgeDate?: Date;
    /**
     * Flag to indicate if the event space is associated from the Package
     * @type {boolean}
     * @memberof EventDetailType
     */
    includeSpaceInPackage?: boolean;
    /**
     * Flag that indicates if actual revenue can be manually edited.
     * @type {boolean}
     * @memberof EventDetailType
     */
    eventLevelRevenueActualization?: boolean;
}

/**
 * Check if a given object implements the EventDetailType interface.
 */
export function instanceOfEventDetailType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function EventDetailTypeFromJSON(json: any): EventDetailType {
    return EventDetailTypeFromJSONTyped(json, false);
}

export function EventDetailTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EventDetailType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'eventName': !exists(json, 'eventName') ? undefined : TranslationTextType60FromJSON(json['eventName']),
        'eventType': !exists(json, 'eventType') ? undefined : json['eventType'],
        'eventTimeSpan': !exists(json, 'eventTimeSpan') ? undefined : DateTimeSpanTypeFromJSON(json['eventTimeSpan']),
        'eventStatus': !exists(json, 'eventStatus') ? undefined : BookingStatusTypeFromJSON(json['eventStatus']),
        'waitlistReturnStatus': !exists(json, 'waitlistReturnStatus') ? undefined : BookingStatusTypeFromJSON(json['waitlistReturnStatus']),
        'cateringStatusType': !exists(json, 'cateringStatusType') ? undefined : CateringStatusTypeTypeFromJSON(json['cateringStatusType']),
        'attendees': !exists(json, 'attendees') ? undefined : CateringEventsAttendeesTypeFromJSON(json['attendees']),
        'masterEventId': !exists(json, 'masterEventId') ? undefined : UniqueIDTypeFromJSON(json['masterEventId']),
        'doorcard': !exists(json, 'doorcard') ? undefined : json['doorcard'],
        'eventLink': !exists(json, 'eventLink') ? undefined : CateringEventLinkTypeFromJSON(json['eventLink']),
        'inactiveDate': !exists(json, 'inactiveDate') ? undefined : (new Date(json['inactiveDate'])),
        'packageId': !exists(json, 'packageId') ? undefined : json['packageId'],
        'packageName': !exists(json, 'packageName') ? undefined : json['packageName'],
        'masterEvent': !exists(json, 'masterEvent') ? undefined : json['masterEvent'],
        'subEvent': !exists(json, 'subEvent') ? undefined : json['subEvent'],
        'packageEvent': !exists(json, 'packageEvent') ? undefined : json['packageEvent'],
        'alternateEvent': !exists(json, 'alternateEvent') ? undefined : json['alternateEvent'],
        'waitlisted': !exists(json, 'waitlisted') ? undefined : json['waitlisted'],
        'notMoveable': !exists(json, 'notMoveable') ? undefined : json['notMoveable'],
        'loudEvent': !exists(json, 'loudEvent') ? undefined : json['loudEvent'],
        'hasPostings': !exists(json, 'hasPostings') ? undefined : json['hasPostings'],
        'displayDoorcard': !exists(json, 'displayDoorcard') ? undefined : json['displayDoorcard'],
        'cateringDeductInventory': !exists(json, 'cateringDeductInventory') ? undefined : json['cateringDeductInventory'],
        'eventShared': !exists(json, 'eventShared') ? undefined : json['eventShared'],
        'createDateTime': !exists(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !exists(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !exists(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !exists(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'purgeDate': !exists(json, 'purgeDate') ? undefined : (new Date(json['purgeDate'])),
        'includeSpaceInPackage': !exists(json, 'includeSpaceInPackage') ? undefined : json['includeSpaceInPackage'],
        'eventLevelRevenueActualization': !exists(json, 'eventLevelRevenueActualization') ? undefined : json['eventLevelRevenueActualization'],
    };
}

export function EventDetailTypeToJSON(value?: EventDetailType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'eventName': TranslationTextType60ToJSON(value.eventName),
        'eventType': value.eventType,
        'eventTimeSpan': DateTimeSpanTypeToJSON(value.eventTimeSpan),
        'eventStatus': BookingStatusTypeToJSON(value.eventStatus),
        'waitlistReturnStatus': BookingStatusTypeToJSON(value.waitlistReturnStatus),
        'cateringStatusType': CateringStatusTypeTypeToJSON(value.cateringStatusType),
        'attendees': CateringEventsAttendeesTypeToJSON(value.attendees),
        'masterEventId': UniqueIDTypeToJSON(value.masterEventId),
        'doorcard': value.doorcard,
        'eventLink': CateringEventLinkTypeToJSON(value.eventLink),
        'inactiveDate': value.inactiveDate === undefined ? undefined : (value.inactiveDate.toISOString().substr(0,10)),
        'packageId': value.packageId,
        'packageName': value.packageName,
        'masterEvent': value.masterEvent,
        'subEvent': value.subEvent,
        'packageEvent': value.packageEvent,
        'alternateEvent': value.alternateEvent,
        'waitlisted': value.waitlisted,
        'notMoveable': value.notMoveable,
        'loudEvent': value.loudEvent,
        'hasPostings': value.hasPostings,
        'displayDoorcard': value.displayDoorcard,
        'cateringDeductInventory': value.cateringDeductInventory,
        'eventShared': value.eventShared,
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'purgeDate': value.purgeDate === undefined ? undefined : (value.purgeDate.toISOString().substr(0,10)),
        'includeSpaceInPackage': value.includeSpaceInPackage,
        'eventLevelRevenueActualization': value.eventLevelRevenueActualization,
    };
}

