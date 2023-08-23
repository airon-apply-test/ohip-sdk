/**
 * OPERA Cloud Integration Configuration API
 * APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ClassOfServiceConfigType } from './ClassOfServiceConfigType';
import type { DataLineType } from './DataLineType';
import type { DirectInwardDialType } from './DirectInwardDialType';
import type { DoNotDisturbType } from './DoNotDisturbType';
import type { MaskDialNumberType } from './MaskDialNumberType';
import type { MessageWaitingType } from './MessageWaitingType';
/**
 *
 * @export
 * @interface InterfaceRoomType
 */
export interface InterfaceRoomType {
    /**
     * Hotel Code of the Interface Room.
     * @type {string}
     * @memberof InterfaceRoomType
     */
    hotelId?: string;
    /**
     * Hotel Interface Logo. On the database, this is also referred as DBF Logo. This is a three letter code followed by an underscore(_). This tells us which DBF files and log files with the prefix that IFC7 is going to create.
     * @type {string}
     * @memberof InterfaceRoomType
     */
    logo?: string;
    /**
     * Front Office Room Type.
     * @type {string}
     * @memberof InterfaceRoomType
     */
    roomType?: string;
    /**
     * Front Office Room Number.
     * @type {string}
     * @memberof InterfaceRoomType
     */
    roomId?: string;
    /**
     * External System's extension number.
     * @type {string}
     * @memberof InterfaceRoomType
     */
    lineNumber?: string;
    /**
     * External System's extension number. It is used for editing a record. It is logically part of the record ID, and can actually be updated.
     * @type {string}
     * @memberof InterfaceRoomType
     */
    oldLineNumber?: string;
    /**
     * Line Type of an Interface Room.
     * @type {string}
     * @memberof InterfaceRoomType
     */
    type?: string;
    /**
     *
     * @type {DataLineType}
     * @memberof InterfaceRoomType
     */
    dataLine?: DataLineType;
    /**
     * Translation Table Id as defined in the IFC_CTRL table. It is also referred as Xlat Id.
     * @type {string}
     * @memberof InterfaceRoomType
     */
    dataLineId?: string;
    /**
     * Value to set/check if the wake up call funtionality is enabled/disabled for the particular line number.
     * @type {boolean}
     * @memberof InterfaceRoomType
     */
    wakeUpCall?: boolean;
    /**
     *
     * @type {DoNotDisturbType}
     * @memberof InterfaceRoomType
     */
    doNotDisturb?: DoNotDisturbType;
    /**
     *
     * @type {ClassOfServiceConfigType}
     * @memberof InterfaceRoomType
     */
    classOfService?: ClassOfServiceConfigType;
    /**
     *
     * @type {DirectInwardDialType}
     * @memberof InterfaceRoomType
     */
    directInwardDial?: DirectInwardDialType;
    /**
     *
     * @type {MessageWaitingType}
     * @memberof InterfaceRoomType
     */
    messageWaiting?: MessageWaitingType;
    /**
     *
     * @type {MaskDialNumberType}
     * @memberof InterfaceRoomType
     */
    maskDialNumber?: MaskDialNumberType;
}
/**
 * Check if a given object implements the InterfaceRoomType interface.
 */
export declare function instanceOfInterfaceRoomType(value: object): boolean;
export declare function InterfaceRoomTypeFromJSON(json: any): InterfaceRoomType;
export declare function InterfaceRoomTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): InterfaceRoomType;
export declare function InterfaceRoomTypeToJSON(value?: InterfaceRoomType | null): any;
