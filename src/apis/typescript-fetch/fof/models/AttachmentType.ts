/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { DateTimeStampGroupType } from './DateTimeStampGroupType';
import {
    DateTimeStampGroupTypeFromJSON,
    DateTimeStampGroupTypeFromJSONTyped,
    DateTimeStampGroupTypeToJSON,
} from './DateTimeStampGroupType';

/**
 * Attached files.
 * @export
 * @interface AttachmentType
 */
export interface AttachmentType {
    /**
     * Name of the file.
     * @type {string}
     * @memberof AttachmentType
     */
    fileName?: string;
    /**
     * Size of the file.
     * @type {number}
     * @memberof AttachmentType
     */
    fileSize?: number;
    /**
     * Description for the file.
     * @type {string}
     * @memberof AttachmentType
     */
    description?: string;
    /**
     * Flag to say if attachment is available across properties.
     * @type {boolean}
     * @memberof AttachmentType
     */
    global?: boolean;
    /**
     * Hotel Code.
     * @type {string}
     * @memberof AttachmentType
     */
    hotelId?: string;
    /**
     * 
     * @type {DateTimeStampGroupType}
     * @memberof AttachmentType
     */
    history?: DateTimeStampGroupType;
    /**
     * URL that identifies the location associated with the record identified by the UniqueID.
     * @type {string}
     * @memberof AttachmentType
     */
    url?: string;
    /**
     * A reference to the type of object defined by the UniqueID element. Refer to OpenTravel Code List Unique ID Type (UIT).
     * @type {string}
     * @memberof AttachmentType
     */
    type?: string;
    /**
     * The identification of a record as it exists at a point in time. An instance is used in update messages where the sender must assure the server that the update sent refers to the most recent modification level of the object being updated.
     * @type {string}
     * @memberof AttachmentType
     */
    instance?: string;
    /**
     * Used to identify the source of the identifier (e.g., IATA, ABTA).
     * @type {string}
     * @memberof AttachmentType
     */
    idContext?: string;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof AttachmentType
     */
    id?: string;
    /**
     * Additional identifying value assigned by the creating system.
     * @type {number}
     * @memberof AttachmentType
     */
    idExtension?: number;
}

/**
 * Check if a given object implements the AttachmentType interface.
 */
export function instanceOfAttachmentType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AttachmentTypeFromJSON(json: any): AttachmentType {
    return AttachmentTypeFromJSONTyped(json, false);
}

export function AttachmentTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AttachmentType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'fileName': !exists(json, 'fileName') ? undefined : json['fileName'],
        'fileSize': !exists(json, 'fileSize') ? undefined : json['fileSize'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'global': !exists(json, 'global') ? undefined : json['global'],
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'history': !exists(json, 'history') ? undefined : DateTimeStampGroupTypeFromJSON(json['history']),
        'url': !exists(json, 'url') ? undefined : json['url'],
        'type': !exists(json, 'type') ? undefined : json['type'],
        'instance': !exists(json, 'instance') ? undefined : json['instance'],
        'idContext': !exists(json, 'idContext') ? undefined : json['idContext'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'idExtension': !exists(json, 'idExtension') ? undefined : json['idExtension'],
    };
}

export function AttachmentTypeToJSON(value?: AttachmentType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'fileName': value.fileName,
        'fileSize': value.fileSize,
        'description': value.description,
        'global': value.global,
        'hotelId': value.hotelId,
        'history': DateTimeStampGroupTypeToJSON(value.history),
        'url': value.url,
        'type': value.type,
        'instance': value.instance,
        'idContext': value.idContext,
        'id': value.id,
        'idExtension': value.idExtension,
    };
}

